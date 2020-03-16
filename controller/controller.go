package controller

import (
	v1 "CapPrice/controller/v1"
	"github.com/gin-gonic/gin"
)

func InitController() *gin.Engine {
	route := gin.Default()

	route.GET("/api/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": true,
		})
	})

	authGroup := route.Group("api/auth")
	{
		authGroup.POST("/register", RegisterUser)
		authGroup.POST("/login", LoginHandler)
		authGroup.POST("/check", CheckTokenHandler)
	}

	smsGroup := route.Group("api/sms")
	{
		smsGroup.POST("fetch", SmsFetchHandler)
		smsGroup.POST("register", SmsRegisterHandler)
		smsGroup.POST("login", SmsLoginHandler)
	}

	apiV1 := route.Group("api/v1")
	{
		v1.RegisterApplicationController(apiV1)
	}

	return route
}
