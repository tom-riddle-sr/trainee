package login

import (
	"database/sql"
	"fmt"
	"trainee/fibertrainee2/repository/accountdata"

	"github.com/dgrijalva/jwt-go" //Go 語言的 JWT 庫。它提供了創建、簽名、驗證和解析 JWT 的功能
	"github.com/gofiber/fiber/v2"
)

func Login(account, password string, DB *sql.DB, c *fiber.Ctx) {
	if err := accountdata.ValidateAccount(account, password, DB); err != nil {
		panic(err)
	} else {
		// 新jwt 指定簽名方法
		token := jwt.New(jwt.SigningMethodHS256)
		// 將 token 的 Claims 轉換為 jwt.MapClaims
		claims := token.Claims.(jwt.MapClaims)
		claims["account"] = account
		claims["blab"] = "河馬吃西瓜"
		// 將 token 簽名並轉換為一個字串
		tokenString, tokenErr := token.SignedString([]byte("hippopotamus eat watermelon"))
		if tokenErr != nil {
			fmt.Println("簽署token失敗")
			panic(err)
		}
		fmt.Println("簽署token成功")

		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = tokenString
		c.Cookie(cookie)
		fmt.Println("存進cookie成功")

	}
}
