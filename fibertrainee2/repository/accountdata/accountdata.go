package accountdata

import (
	"database/sql"
	"trainee/fibertrainee2/model"
)

func GetAccount(account string, db *sql.DB) (model.AccountData, error) {

	// sqlStatement := fmt.Sprintf(`SELECT * FROM accountdata WHERE account='%s' AND password='%s'`, hashstring.HashString(account), hashstring.HashString(password))
	// row := db.QueryRow(sqlStatement)
	// if err := row.Scan(&account, &password); err != nil {
	// 	fmt.Println("查詢資料失敗", err)
	// 	return err
	// }
	// fmt.Println("查詢資料成功")
	// return nil

	//直接寫在QueryRow裡面是為了避免sql injection
	var ad model.AccountData

	if err := db.QueryRow(
		`SELECT * FROM accountdata WHERE account='%s'`,
		account,
	).Scan(&ad.Account, &ad.Password); err != nil {
		// sql.ErrNoRows 如果查詢不到資料就回傳空的model.AccountData{}和nil
		if err == sql.ErrNoRows {
			return model.AccountData{}, nil
		}
		return model.AccountData{}, err
	}
	return ad, nil
}
