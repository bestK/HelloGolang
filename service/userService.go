package service

import (
	"hello/model"
	"github.com/donnie4w/go-logger/logger"
	"hello/util"
	"database/sql"
	"github.com/labstack/echo"
)

func SelectUsers(c echo.Context) (*[]model.User, error) {

	curPage := c.QueryParam("curPage")
	pageSize := c.QueryParam("pageSize")

	if curPage == "" || pageSize == "" {
		return nil, util.Err("parameter invalid")
	}

	user := []model.User{}
	selectSql := "SELECT id,nick_name FROM vino_user LIMIT  ?,?"
	if err := db.Select(&user, selectSql, curPage, pageSize); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return &user, nil
}
