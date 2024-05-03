package login

import (
	"database/sql"
	"fmt"
	"trainee/fibertrainee2/repository/accountdata"

	"trainee/fibertrainee2/tools/jwtToken"

	"github.com/gofiber/fiber/v2"
)

func Login(account, password string, DB *sql.DB, c *fiber.Ctx) error {
	if err := accountdata.ValidateAccount(account, password, DB); err != nil {
		fmt.Println("驗證帳密失敗")
		return err
	} else {
		claims := map[string]interface{}{
			"account": "myAccount",
			"blab":    "河馬吃西瓜",
		}
		const (
			signature string = "hippopotamus eat watermelon"
		)
		tokenString, err := jwtToken.JwtToken(claims, signature)
		if err != nil {
			fmt.Println("簽署token失敗")
			return err
		}

		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = tokenString
		c.Cookie(cookie)
		fmt.Println("存進cookie成功")
	}
	return nil
}
