package repository

import (
	"database/sql"
	"trainee/fibertrainee3/model"
)

func Query(db *sql.DB, inputAccountData model.AccountData) (model.AccountData, error) {
	row := db.QueryRow("SELECT * FROM accountdata WHERE account = ? ", inputAccountData.Account)
	queryData := model.AccountData{}
	if err := row.Scan(&queryData.Account, &queryData.Password); err != nil {
		return model.AccountData{}, err
	}
	return queryData, nil
}
