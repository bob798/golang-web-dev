package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func Db()  *gorm.DB {
	url := "root:chbigdata@tcp(139.155.77.236:23346)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}