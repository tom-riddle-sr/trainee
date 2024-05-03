package jwtToken

import (
	"github.com/dgrijalva/jwt-go" //Go 語言的 JWT 庫。它提供了創建、簽名、驗證和解析 JWT 的功能)
)

func JwtToken(claims map[string]interface{}, signature string) (string, error) {
	// 新jwt 指定簽名方法
	token := jwt.New(jwt.SigningMethodHS256)
	// 將 token 的 Claims 轉換為 jwt.MapClaims
	mapClaims := token.Claims.(jwt.MapClaims)
	for key, value := range claims {
		mapClaims[key] = value
	}
	// 將 token 簽名並轉換為一個字串
	tokenString, tokenErr := token.SignedString([]byte(signature))
	if tokenErr != nil {
		return "", tokenErr
	}
	return tokenString, nil
}
