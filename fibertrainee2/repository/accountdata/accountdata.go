package accountdata

import (
	"database/sql"
	"fmt"
	"trainee/fibertrainee2/tools/hashstring"
)

func ValidateAccount(account, password string, DB *sql.DB) error {
	sqlStatement := fmt.Sprintf(`SELECT * FROM accountdata WHERE account='%s' AND password='%s'`, hashstring.HashString(account), hashstring.HashString(password))
	row := DB.QueryRow(sqlStatement)
	if err := row.Scan(&account, &password); err != nil {
		fmt.Println("查詢資料失敗", err)
		return err
	}
	fmt.Println("查詢資料成功")
	return nil
}
