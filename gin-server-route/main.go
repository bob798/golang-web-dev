package main

import (
	"gin-server-route/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route :=gin.Default()

	// 以下测试接口
	// 访问地址 http://localhost:8081/health
	route.GET("/health", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"ok"})
	})

	// 访问地址 http://localhost:8081/test/bob
	route.GET("/test/:name", func(c *gin.Context) {
		name :=c.Param("name")
		c.String(http.StatusOK,"Hello %s",name)
	})

	// 以下正式使用接口
	//注册监控检测路由
	router.InitHealthRouter(route.Group(""))
	//注册用户模块路由
	router.InitUserRouter(route.Group(""))
	route.Run(":8081")
}
