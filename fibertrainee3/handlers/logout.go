package handlers

import (
	"trainee/fibertrainee3/model"
	"trainee/fibertrainee3/tools/cookie"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	token := cookie.Get(c, cookie.TokenName)
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.CommonResponse{
			Err: "token不存在",
		})
	}
	c.ClearCookie(cookie.TokenName)
	return c.SendString("Logout成功")
}
