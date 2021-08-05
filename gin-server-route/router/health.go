package router

import (
	"gin-server-route/api"
	"github.com/gin-gonic/gin"
)

func InitHealthRouter(router *gin.RouterGroup) {
	health := router.Group("health")
	{
		// 访问地址：http://localhost:8081/health/info
		health.GET("info", api.Health)
	}

}
