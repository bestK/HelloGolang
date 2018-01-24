package model

import (
	"hello/util"
)

type User struct {
	Id       util.NullInt64  `db:"id" json:"id"`
	NickName util.NullString  `db:"nick_name" json:"nickName"`
}
