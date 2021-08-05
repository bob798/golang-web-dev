package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router :=gin.Default()
	router.GET("/health", func(c *gin.Context) {
		// response json  data: "message":"success"
		//var data = map[string]interface{}
		//data["message"] = "success"
		c.JSON(200,gin.H{"username":"bob"})
		//c.String(200,"ok")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		// 获取name参数值
		name :=c.Param("name")
		if name =="好人" {
			c.String(http.StatusOK,"Hello %s",name)
		}else {
			// 返回前端数据
			c.String(http.StatusBadRequest,"Hello %s",name)
		}
	})

	// /user/list ， /user/name
	router.GET("/user/*action", func(c *gin.Context) {
		// 获取name参数值
		name :=c.Param("name")
		if name =="人民警察" {
			c.String(http.StatusOK,"Hello %s",name)
		}else {
			// 返回前端数据
			c.String(http.StatusBadRequest,"Hello %s",name)
		}
	})

	router.Run(":8081")
}
