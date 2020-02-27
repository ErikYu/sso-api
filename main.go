package main

import (
	"CapPrice/base"
	"CapPrice/controller"
	"CapPrice/logging"
	"CapPrice/model"
)

func main() {
	err := base.InitConfig()
	if err != nil {
		logging.STDInfo("读取配置文件错误: %v", err)
	}
	model.InitDb()

	route := controller.InitController()

	route.Run(":9999")
}
