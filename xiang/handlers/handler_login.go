package handlers

import (
	"fmt"
	"trainee/fibertrainee2/model"
	"trainee/fibertrainee2/service/login"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	data := model.AccountData{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data.Account == "" || data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "帳密不得為空",
		})
	}
	fmt.Println("驗證參數成功")
	tokenString, err := login.Login(data.Account, data.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: tokenString,
	})

	return c.SendString("登入成功")
}

// TODO: 看一下
func Logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.SendString("登出成功")
}
