package service

import (
	"gin-server-route/config"
	"gin-server-route/model"
)

func ListUser() interface{} {
	db := config.Db()
	var users []model.User
	 db.Find(&users)
	return users
}
