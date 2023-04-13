package dao

import (
	"fmt"
	"github.com/pigeon/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	username := config.GetConfig().MySQL.User     //账号
	password := config.GetConfig().MySQL.Password //密码
	host := config.GetConfig().MySQL.Host         //数据库地址，可以是Ip或者域名
	port := config.GetConfig().MySQL.Port         //数据库端口
	Dbname := config.GetConfig().MySQL.Name       //数据库名
	timeout := "10s"                              //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = _db
}

func GetDB() *gorm.DB {
	return db
}
