package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/labstack/echo"
	"hello/util"
	"net/http"
)

type User struct {
	Name sql.NullString `json:"name" xml:"name" form:"name" query:"name"`
	Id   sql.NullInt64  `json:"Id" xml:"Id" form:"Id" query:"Id"`
}

var (
	db = &sql.DB{}
)

/*
	初始化数据库连接
*/
func init() {
	db, _ = sql.Open("mysql", "root:vinohobby@tcp(127.0.0.1:3306)/vino?charset=utf8")
}

/*
	数据库查询
*/
func selectUser(c echo.Context) error {
	curPage := c.QueryParam("curPage")
	pageSize := c.QueryParam("pageSize")

	if curPage == "" || pageSize == "" {
		return c.JSON(http.StatusBadRequest, util.Fr(http.StatusBadRequest, "parameter invalid", nil))
	}

	return c.JSON(http.StatusOK, util.Fr(http.StatusOK, "successfully", util.Querys("SELECT id,nick_name FROM vino_user LIMIT "+curPage+" , "+pageSize, db)))
}

func main() {
	e := echo.New()
	e.GET("/", selectUser)
	e.Logger.Fatal(e.Start(":1323"))
}
