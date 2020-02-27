package model

import (
	"CapPrice/logging"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDb() {
	var err error
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=localhost port=%s user=%s dbname=%s password=%s sslmode=disable",
			viper.GetString("db_port"),
			viper.GetString("db_user"),
			viper.GetString("db_name"),
			viper.GetString("db_password"),
		),
	)
	if err != nil {
		logging.STDError("连接数据库失败: %v", err)
	}
	db.AutoMigrate(&SsoUser{}, &SsoApp{}, &SsoUserAppRel{})
}
