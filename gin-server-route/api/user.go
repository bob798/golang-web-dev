package api

import (
	"gin-server-route/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"data":service.ListUser()})
}
