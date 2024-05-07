package repository

import (
	"database/sql"
	"errors"
	"trainee/fibertrainee3/tools"
)

const (
	insertAccount  string = "test"
	insertPassword string = "test"
)

func Insert(db *sql.DB) error {

	hashPassword := tools.Sha512(insertPassword)
	if _, err := db.Exec(
		`INSERT INTO accountdata (account, password) VALUES (?, ?)`, insertAccount, hashPassword,
	); err != nil {
		return errors.New("insert 資料失敗")
	}
	return nil
}
