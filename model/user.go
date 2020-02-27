package model

import (
	"CapPrice/logging"
	"crypto/sha512"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"io"
	"math/rand"
	"strings"
)

type SsoUser struct {
	BaseModel
	Login        string `gorm:"UNIQUE"`
	PwdEncrypted string
	Salt         string
	Cellphone    string `gorm:"UNIQUE"`
}

func (user SsoUser) validate(password string) bool {
	pwdEncrypted := encrypt(password, user.Salt)
	return pwdEncrypted == user.PwdEncrypted
}

func (user SsoUser) generateJwt() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": user.Login,
	})
	SECRET := viper.GetString("crendentials.jwt_secret")
	tokenString, _ := token.SignedString([]byte(SECRET))
	return tokenString
}

func GetUserByLogin(login string) (user SsoUser) {
	db.Where(&SsoUser{
		Login: login,
	}).First(&user)
	return
}

func GetUserByCellphone(cellphone string) (user SsoUser) {
	db.Where(&SsoUser{
		Cellphone: cellphone,
	}).First(&user)
	return
}

type SsoUserAppRel struct {
	BaseModelOnlyId
	UserId uint
	AppId  uint
}

func getRandomStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func encrypt(pwd, salt string) string {
	hashed := sha512.New512_256()
	_, _ = io.WriteString(hashed, strings.Join([]string{pwd, salt}, ""))
	return fmt.Sprintf("%x", hashed.Sum(nil))
}

func CreateUserByLogin(login, password, appName string) error {

	// get app by appName
	app := GetAppByName(appName)
	if app.ID == 0 {
		return fmt.Errorf("未找到该App -> %s", appName)
	}

	// encrypt the password
	salt := getRandomStr(8)
	pwdEncrypted := encrypt(password, salt)

	// insert user and user_app_rel
	return db.Transaction(func(tx *gorm.DB) error {
		newUser := SsoUser{
			Login:        login,
			PwdEncrypted: pwdEncrypted,
			Salt:         salt,
			Cellphone:    "",
		}
		res := db.Create(&newUser)
		if res.Error != nil {
			return res.Error
		}
		newUserAppRel := SsoUserAppRel{
			UserId: res.Value.(*SsoUser).ID,
			AppId:  app.ID,
		}
		createRelMutation := db.Create(&newUserAppRel)

		if createRelMutation.Error != nil {
			logging.STDError("创建用户应用关联错误: %v", createRelMutation.Error)
			return createRelMutation.Error
		}
		return nil
	})
}

func ValidateUserByLogin(login, password string) (string, error) {
	currentUser := GetUserByLogin(login)
	if currentUser.ID == 0 {
		return "", fmt.Errorf("未找到当前用户: %s", login)
	}
	if !currentUser.validate(password) {
		return "", fmt.Errorf("密码错误")
	}
	return currentUser.generateJwt(), nil
}

func CreateUserByCellphone(cellphone, appName string) error {
	// get app by appName
	app := GetAppByName(appName)
	if app.ID == 0 {
		return fmt.Errorf("未找到该App -> %s", appName)
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// make the cellphone as login default
		newUser := SsoUser{
			Login:        fmt.Sprintf("手机用户_%s", cellphone),
			PwdEncrypted: "",
			Salt:         "",
			Cellphone:    cellphone,
		}
		createUserMutation := tx.Create(&newUser)
		if createUserMutation.Error != nil {
			return createUserMutation.Error
		}

		newUserAppRel := SsoUserAppRel{
			UserId: createUserMutation.Value.(*SsoUser).ID,
			AppId:  app.ID,
		}
		createRelMutation := db.Create(&newUserAppRel)

		if createRelMutation.Error != nil {
			logging.STDError("创建用户应用关联错误: %v", createRelMutation.Error)
			return createRelMutation.Error
		}
		return nil
	})
}
