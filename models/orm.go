package models

import (
	"../conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strconv"
	"time"
)

var orm *gorm.DB

// func ConnectDB() {

// }

func init() {
	// conf.GetConfig("mysql.username")
	var err error
	// 主机
	host := conf.GetConfig("mysql.host")
	// 端口
	port := conf.GetConfig("mysql.port")
	// 用户名
	username := conf.GetConfig("mysql.username")
	// 密码
	password := conf.GetConfig("mysql.password")
	// 数据库名称
	dbname := conf.GetConfig("mysql.dbname")
	// 最大空闲连接数
	maxIdleConns, _ := strconv.Atoi(conf.GetConfig("mysql.max_idle_conns"))
	// 最大打开的连接数
	maxOpenConns, _ := strconv.Atoi(conf.GetConfig("mysql.max_open_conns"))

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"

	if orm, err = gorm.Open("mysql", dsn); err != nil {
		fmt.Printf("Fail to open mysql: %v", err)
		os.Exit(1)
	}

	orm.DB().SetMaxIdleConns(maxIdleConns)
	orm.DB().SetMaxOpenConns(maxOpenConns)
	orm.DB().SetConnMaxLifetime(5 * time.Minute)
}
