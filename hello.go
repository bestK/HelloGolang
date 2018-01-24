package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/labstack/echo"
	"hello/controller"
	"github.com/jmoiron/sqlx"
	"hello/service"
)

type User struct {
	Name sql.NullString `json:"name" xml:"name" form:"name" query:"name"`
	Id   sql.NullInt64  `json:"Id" xml:"Id" form:"Id" query:"Id"`
}

var (
	db *sqlx.DB
)


func main() {

	// 连接数据库
	service.Start()



	e := echo.New()
	e.GET("/api/users", controller.SelectUser)
	//e.GET("/", index)
	e.Logger.Fatal(e.Start(":80"))
}
