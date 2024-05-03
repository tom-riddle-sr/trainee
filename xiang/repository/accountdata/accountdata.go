package accountdata

import (
	"database/sql"

	"trainee/fibertrainee2/model"
)

func GetAccount(account string, db *sql.DB) (model.AccountData, error) {
	var ad model.AccountData
	if err := db.QueryRow(
		`SELECT * FROM accountdata WHERE account='%s'`,
		account,
	).Scan(&ad.Account, &ad.Password); err != nil {
		if err == sql.ErrNoRows {
			return model.AccountData{}, nil
		}
		return model.AccountData{}, err
	}
	return ad, nil
}
