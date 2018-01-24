package service

import (
	"github.com/donnie4w/go-logger/logger"
	"github.com/jmoiron/sqlx"
)

var (
	//数据库操作对象
	db *sqlx.DB
)

/**
	初始化数据库连接
 */
func Start() {
	logger.Info("Start connection databases...")
	sqlSession, err := sqlx.Open("mysql", "root:vinohobby@tcp(127.0.0.1:3306)/vino?charset=utf8")
	if err != nil {
		panic("Open database error!")
	}

	if err := sqlSession.Ping(); err != nil {
		panic("Connection database fail!")
	}

	db = sqlSession
	logger.Info("Connection database success...")
}
