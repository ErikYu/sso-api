package controller

import (
	"CapPrice/logging"
	"CapPrice/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type SmsPayload struct {
	Cellphone string `json:"cellphone"`
}

type SmsCodeDetail struct {
	Code     string
	ExpireAt time.Time
}

var UserSmsCodeMap = map[string][]SmsCodeDetail{}

func SmsFetchHandler(context *gin.Context) {
	var payload SmsPayload
	err := context.BindJSON(&payload)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "获取手机号失败",
		})
		return
	}
	// create verifyCode
	letters := []rune("1234567890")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	code := string(b)

	logging.STDDebug("验证码FAKE: %s", code)

	newCode := SmsCodeDetail{
		Code:     code,
		ExpireAt: time.Now().Add(5 * time.Minute),
	}
	UserSmsCodeMap[payload.Cellphone] = append(UserSmsCodeMap[payload.Cellphone], newCode)
	// fetch the sms code through aliyun
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	return
}

type SmsValidatePayload struct {
	Cellphone   string `json:"cellphone"`
	VerifyCode  string `json:"verifyCode"`
	Application string `json:"application"`
}

// Use sms code to register
func SmsRegisterHandler(context *gin.Context) {
	var payload SmsValidatePayload
	err := context.BindJSON(&payload)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	// check if verifyCode correct
	if !checkIfVerifyCodeValid(payload.Cellphone, payload.VerifyCode) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "验证码错误",
		})
		return
	}

	// check if cellphone exists
	user := model.GetUserByCellphone(payload.Cellphone)
	if user.ID != 0 {
		logging.STDError("验证码注册账号失败，该手机号已存在: %s", payload.Cellphone)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "验证码注册账号失败，该手机号已存在",
		})
		return
	}

	// create user
	err = model.CreateUserByCellphone(payload.Cellphone, payload.Application)
	if err != nil {
		logging.STDError("验证码注册用户失败: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "验证码注册用户失败",
			"trace":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func SmsValidateHandler(context *gin.Context) {
	var payload SmsValidatePayload
	err := context.BindJSON(&payload)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	var isCodeValidate = false
	for _, codeDetail := range UserSmsCodeMap[payload.Cellphone] {
		fmt.Println(codeDetail.Code)
		if codeDetail.Code == payload.VerifyCode && time.Now().Before(codeDetail.ExpireAt) {
			isCodeValidate = true
		}
	}
	if !isCodeValidate {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "验证码错误",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func checkIfVerifyCodeValid(cellphone, code string) (isCodeValidate bool) {
	for _, codeDetail := range UserSmsCodeMap[cellphone] {
		fmt.Println(codeDetail.Code)
		if codeDetail.Code == code && time.Now().Before(codeDetail.ExpireAt) {
			isCodeValidate = true
		}
	}
	return
}
