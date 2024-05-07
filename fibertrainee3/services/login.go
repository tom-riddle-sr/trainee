package services

import (
	"database/sql"
	"errors"
	"fmt"
	"trainee/fibertrainee3/database/mysql"
	"trainee/fibertrainee3/model"
	"trainee/fibertrainee3/repository"
	"trainee/fibertrainee3/tools"
	"trainee/fibertrainee3/tools/jwt"
)

func Login(inputAccountData model.AccountData) (string, error) {
	queryData, err := repository.Query(mysql.GetDB(), inputAccountData)
	fmt.Println(queryData)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("查無此帳號")
		}
		return "", fmt.Errorf("查詢資料庫錯誤: %w", err)
	}

	if queryData.Password != tools.Sha512(inputAccountData.Password) {
		return "", errors.New("密碼錯誤")
	}
	encryptData, err := jwt.GenToken(queryData, jwt.Secret)
	if err != nil {
		return "", fmt.Errorf("加密錯誤: %w", err)
	}
	return encryptData, nil
}
