package router

import (
	"github.com/labstack/echo"
	"hello/controller"
	"github.com/donnie4w/go-logger/logger"
)

/*
	请求路由
*/
func Start()  {
	e := echo.New()
	e.GET("/api/users", controller.SelectUser)
	e.GET("/", controller.Index)
	e.Logger.Fatal(e.Start(":80"))
	logger.Info("Router build done")
}
