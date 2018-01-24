package main

import (
	"hello/router"
	"hello/service"
)

func main() {

	// 连接数据库
	service.Start()

	// 配置路由
	router.Start()

}
