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
func SelectUser_(c echo.Context) error {
	curPage := c.QueryParam("curPage")
	pageSize := c.QueryParam("pageSize")

	if curPage == "" || pageSize == "" {
		return c.JSON(http.StatusBadRequest, util.Fr(http.StatusBadRequest, "parameter invalid", nil))
	}

	return c.JSON(http.StatusOK, util.Fr(http.StatusOK, "successfully", util.Querys("SELECT id,nick_name FROM vino_user LIMIT "+curPage+" , "+pageSize, nil)))
}

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
