package router

import (
	"gin-server-route/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	health := router.Group("user")
	{
		// 接口地址：http://localhost:8081/user/getUserList
		health.GET("getUserList", api.GetUserList)
	}

}
