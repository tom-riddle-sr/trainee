package main

import (
	"database/sql"
	"fmt"
)

func Create(DB *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS Member2(
    MemberId INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Status CHAR(1) NOT NULL,
    CreateTime DATETIME NOT NULL,
    UpdateTime DATETIME NOT NULL
	); `

	if _, err := DB.Exec(sql)
	err != nil {
		fmt.Println("創建新table失敗", err)
		return
	}
	fmt.Println("創建新table成功")
}