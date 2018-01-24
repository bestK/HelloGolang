package controller

import (
	"github.com/labstack/echo"
	"hello/service"
)

/*
	数据库查询
*/
func SelectUser(c echo.Context) error {
	return service.SelectUsers(c)
}
