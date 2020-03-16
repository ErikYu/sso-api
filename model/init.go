package model

import (
	"CapPrice/logging"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDb() {
	var (
		//ServerMode = viper.GetString("server.mode")
		DBPort     = viper.GetString("db_port")
		DBUser     = viper.GetString("db_user")
		DBName     = viper.GetString("db_name")
		DBPassword = viper.GetString("db_password")
	)
	var err error
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=localhost port=%s user=%s dbname=%s password=%s sslmode=disable",
			DBPort,
			DBUser,
			DBName,
			DBPassword,
		),
	)
	if err != nil {
		logging.STDError("连接数据库失败: %v", err)
	}
	db.AutoMigrate(&SsoUser{}, &SsoApp{}, &SsoUserAppRel{})
}
