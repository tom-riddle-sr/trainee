package token

import (
	"fmt"
	"trainee/fibertrainee/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// what is jwt(JSON Web Token)?
// 將資訊（"claims"）包裝成 JSON 物件的方式，並將其編成一個字串
// 這個字串可以在不同的 URL 或者 HTTP headers 中傳遞

// Header
// Payload
// Signature 由 Header、Payload 和一個秘密鍵使用簽名算法生成

// fiber.Handler 是類型別名
// 這個函數接收一個 *fiber.Ctx 參數，並返回一個 error

func Token() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 新jwt 指定簽名方法
		token := jwt.New(jwt.SigningMethodHS256)
		// 將 token 的 Claims 轉換為 jwt.MapClaims
		claims := token.Claims.(jwt.MapClaims)
		claims["account"] = common.Postdata["account"]
		claims["blab"] = "河馬吃西瓜"
		// 將 token 簽名並轉換為一個字串
		tokenString, err := token.SignedString([]byte("hippopotamus eat watermelon"))
		if err != nil {
			fmt.Println("簽署token失敗")
			return err
		}
		// 將 token 存儲在 *fiber.Ctx 中,可以在後續中間件中使用
		c.Locals("token", tokenString)
		fmt.Println("簽署成功 Token:")

		return nil
	}
}
