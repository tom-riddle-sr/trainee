package mysql

import (
	"database/sql"
	"fmt"
)

const (
	UserName = "root"
	Password = "yile1408"
	Addr     = "127.0.0.1"
	Port     = 3003
	Database = "fibertrainee"
)

var _db *sql.DB

func New() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		UserName, Password, Addr, Port, Database))
	if err != nil {
		fmt.Println("連接到mysql 失敗:", err)
		return
	}
	_db = db
}

func GetDB() *sql.DB {
	return _db
}
