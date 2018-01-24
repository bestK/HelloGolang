package controller

import (
	"hello/util"
	"github.com/labstack/echo"
	"net/http"
	"hello/service"
)

/*
	数据库查询
*/
func SelectUser(c echo.Context) error {

	data, err := service.SelectUsers(c)

	if err != nil {
		return util.Bad(c, http.StatusInternalServerError, err.Error())
	}

	return util.Ok(c, "successfully", data)
}
