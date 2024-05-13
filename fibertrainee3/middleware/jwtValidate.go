package middleware

import (
	"trainee/fibertrainee3/tools/cookie"
	"trainee/fibertrainee3/tools/jwt"

	"trainee/fibertrainee3/model/output"

	"github.com/gofiber/fiber/v2"
)

func JwtValidate(c *fiber.Ctx) error {
	token := cookie.Get(c, cookie.TokenName)
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: "token 不存在",
		})
	}
	if err := jwt.ValidateToken(token, jwt.Secret); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: err.Error(),
		})
	}
	return c.Next()
}
