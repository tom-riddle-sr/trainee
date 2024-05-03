package login

import (
	"errors"
	"fmt"
	"trainee/fibertrainee2/database/mySql"
	"trainee/fibertrainee2/repository/accountdata"

	hashstring "trainee/fibertrainee2/tools/hashString"
	"trainee/fibertrainee2/tools/jwtToken"
)

const (
	signature string = "hippopotamus eat watermelon"
)

func Login(account, password string) (string, error) {
	switch data, err := accountdata.GetAccount(account, mySql.GetDB()); {
	case err != nil:
		return "", err
	case data.Account == "":
		return "", errors.New("帳號錯誤")
	case data.Password != hashstring.HashString(password):
		return "", errors.New("密碼錯誤")
	}
	claims := map[string]interface{}{
		"account": "myAccount",
		"blab":    "河馬吃西瓜",
	}
	tokenString, err := jwtToken.JwtToken(claims, signature)
	if err != nil {
		fmt.Println("簽署token失敗")
		return "", err
	}
	fmt.Println("簽署toke成功")
	return tokenString, nil
}

//為何要在service就把db當參數傳進去,而不是在repository就直接取得db?
// 有可能要做交易transaction,會有多個repository不一定要用同一個db
