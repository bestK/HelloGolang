package service

import (
	"hello/model"
	"github.com/donnie4w/go-logger/logger"
	"hello/util"
	"github.com/labstack/echo"
	"net/http"
)

func SelectUsers(c echo.Context) error {
	curPage := c.QueryParam("curPage")
	pageSize := c.QueryParam("pageSize")

	if curPage == "" || pageSize == "" {
		return util.Bad(c, http.StatusBadRequest, "parameter invalid")
	}

	var user []model.User
	selectSql := "SELECT id,nick_name FROM vino_user LIMIT  ?,?"
	if err := db.Select(&user, selectSql, curPage, pageSize); err != nil {
		logger.Error(err)
		return util.Bad(c, http.StatusInternalServerError, err.Error())
	}

	return util.Ok(c, "successfully", &user)
}
