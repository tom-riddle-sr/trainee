package mysql

import (
	"fmt"
	"trainee/fibertrainee3/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

// Mysql版
func New() error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		config.UserName, config.Password, config.Addr, config.Port, config.Database)))
	if err != nil {
		fmt.Println("連接到mysql 失敗:", err)
		return err
	}
	_db = db
	fmt.Println("連接到mysql 成功")
	return nil
}

func GetDB() *gorm.DB {
	return _db
}
