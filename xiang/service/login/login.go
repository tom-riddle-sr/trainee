package login

import (
	"errors"
	"fmt"
	"trainee/fibertrainee2/database/mysql"
	"trainee/fibertrainee2/repository/accountdata"

	"trainee/fibertrainee2/tools/hashString"
	"trainee/fibertrainee2/tools/jwtToken"
)

const (
	signature string = "hippopotamus eat watermelon"
)

func Login(account, password string) (string, error) {
	switch ad, err := accountdata.GetAccount(account, mysql.GetDB()); {
	case err != nil:
		return "", err
	case ad.Account == "":
		return "", errors.New("帳號錯誤")
	case ad.Password != hashString.HashString(password):
		return "", errors.New("密碼錯誤")
	}

	claims := map[string]interface{}{
		"account": "myAccount",
		"blab":    "河馬吃西瓜",
	}

	tokenString, err := jwtToken.JwtToken(claims, signature)
	if err != nil {
		return "", fmt.Errorf("簽署token失敗: %w", err)
	}

	return tokenString, nil
}
