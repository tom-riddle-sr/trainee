package jwt

import (
	"errors"
	"fmt"

	"trainee/fibertrainee3/model/entity/mysql/fibertrainee"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenName = "loginData"
	Secret    = "bibibaba"
)

func GenToken(value fibertrainee.AccountDatum, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims.(jwt.MapClaims)["account"] = value.Account
	tokenString, tokenErr := token.SignedString([]byte(secret))
	if tokenErr != nil {
		return "", tokenErr
	}
	return tokenString, nil
}

func ValidateToken(tokenString, secret string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return errors.New("token解析錯誤")
	}

	fmt.Println("token解析成功")
	return nil
}
