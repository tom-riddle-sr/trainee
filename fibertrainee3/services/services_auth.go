package services

import (
	"database/sql"
	"errors"
	"fmt"
	"trainee/fibertrainee3/database/mysql"
	"trainee/fibertrainee3/model/entity/mysql/fibertrainee"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/repository"
	"trainee/fibertrainee3/tools"
	"trainee/fibertrainee3/tools/jwt"
)

type IServicesAuth interface {
	Login(inputAccountData input.CreateAccountData) (string, error)
}

type ServicesAuth struct {
	repo *repository.Repo
}

func NewAuth(repo *repository.Repo) IServicesAuth {
	return &ServicesAuth{
		repo: repo,
	}
}

func (h *ServicesAuth) Login(inputAccountData input.CreateAccountData) (string, error) {
	accountdata := fibertrainee.AccountDatum{}

	//型別斷言
	err := h.repo.SqlRepo.Query(mysql.GetDB(), "account = ?", &accountdata, inputAccountData.Account)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("查無此帳號")
		}
		return "", fmt.Errorf("查詢資料庫錯誤: %w", err)
	}

	if accountdata.Password != tools.Sha512(inputAccountData.Password) {
		return "", errors.New("密碼錯誤")
	}
	encryptData, err := jwt.GenToken(accountdata, jwt.Secret)
	if err != nil {
		return "", fmt.Errorf("加密錯誤: %w", err)
	}
	return encryptData, nil
}
