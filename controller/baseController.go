package controller

import (
	"github.com/labstack/echo"
	"hello/util"
	"net/http"
)

func NotFoundHandler(c echo.Context) error {
	return util.Bad(c, http.StatusNotFound, "Page not found")
}

/*
	首页
*/
func Index(c echo.Context) error {
	return util.Ok(c, "service run", nil)
}
