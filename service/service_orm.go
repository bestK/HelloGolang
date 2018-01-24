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
	logger.Info("start connection databases...")
	sqlSession, err := sqlx.Open("mysql", "root:vinohobby@tcp(127.0.0.1:3306)/vino?charset=utf8")
	if err != nil {
		panic("open database error!")
	}

	if err := sqlSession.Ping(); err != nil {
		panic("connection database fail!")
	}

	db = sqlSession
	logger.Info("connection database success...")

	// 测试数据库查询
	var count int
	if err = db.Get(&count, "SELECT count(*) FROM vino_user"); err != nil {
		logger.Error(err)
		return
	}
	logger.Debug("count t_user :", count)

}
