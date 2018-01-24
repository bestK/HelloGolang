package main

import (
	_ "github.com/go-sql-driver/mysql"
	"hello/router"
	"github.com/jmoiron/sqlx"
	"hello/service"
)

var (
	db *sqlx.DB
)


func main() {

	// 连接数据库
	service.Start()

	// 配置路由
	router.Start()

}
