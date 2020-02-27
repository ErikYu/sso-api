package model

import (
	"CapPrice/base/conf"
	"CapPrice/logging"
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDb() {
	var err error
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=localhost port=%s user=%s dbname=%s password=%s sslmode=disable",
			conf.DBPort,
			conf.DBUser,
			conf.DBName,
			conf.DBPassword,
		),
	)
	if err != nil {
		logging.STDError("连接数据库失败: %v", err)
	}
	db.AutoMigrate(&SsoUser{}, &SsoApp{}, &SsoUserAppRel{})
}
