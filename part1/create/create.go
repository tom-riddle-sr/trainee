package create

import (
	"database/sql"
	"fmt"
)

func Create(DB *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS Member(
    MemberId VARCHAR(50) PRIMARY KEY NOT NULL,
    Status CHAR(1) NOT NULL,
    CreateTime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdateTime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	); `

	if _, err := DB.Exec(sql)
	err != nil {
		fmt.Println("創建新table失敗", err)
		return
	}
	fmt.Println("創建新table成功")
}