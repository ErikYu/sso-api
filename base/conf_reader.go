package base

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func InitConfig() error {
	viper.AddConfigPath("./") // 如果没有指定配置文件，则解析默认的配置文件
	viper.SetConfigName("conf.yaml")
	viper.SetConfigType("yaml")                  // 设置配置文件格式为YAML
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		//log.Printf("解析配置文件失败：%v\n", err)
		return err
	}
	return nil
}
