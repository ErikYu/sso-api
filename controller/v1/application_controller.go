package v1

import "github.com/gin-gonic/gin"

func RegisterApplicationController(r *gin.RouterGroup) {
	r.GET("application/list", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": true,
		})
	})
}
