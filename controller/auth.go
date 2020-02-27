package controller

import (
	"CapPrice/logging"
	"CapPrice/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterPayload struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	Application string `json:"application"`
}

func RegisterUser(context *gin.Context) {
	var payload RegisterPayload
	err := context.BindJSON(&payload)
	if err != nil {
		logging.STDError("获取注册参数失败")
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "获取注册参数失败",
		})
		return
	}
	err = model.CreateUserByLogin(payload.Login, payload.Password, payload.Application)
	if err != nil {
		logging.STDError("创建用户失败: %v", err)
		context.JSON(500, gin.H{
			"message": "创建用户失败",
			"trace":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	return
}

func LoginHandler(context *gin.Context) {
	var payload RegisterPayload
	err := context.BindJSON(&payload)
	if err != nil {
		logging.STDError("获取登录参数失败")
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "获取登录参数失败",
		})
		return
	}
	jwt, err := model.ValidateUserByLogin(payload.Login, payload.Password)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "登录失败",
			"trace":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    jwt,
	})
}
