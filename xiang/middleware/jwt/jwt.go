package jwt

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func New() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("錯誤,沒有token存在")
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "default-secret" // 如果環境變量中沒有秘密，則使用默認秘密
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).SendString("錯誤,token無效")
		}

		return c.Next()
	}
}
