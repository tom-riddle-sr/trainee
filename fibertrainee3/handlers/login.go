package handlers

import (
	"fmt"
	"trainee/fibertrainee3/model"
	"trainee/fibertrainee3/services"
	"trainee/fibertrainee3/tools/cookie"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	inputAccountData := model.AccountData{}
	if err := c.BodyParser(&inputAccountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.CommonResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	if inputAccountData.Account == "" || inputAccountData.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			model.CommonResponse{
				Err: "帳號或密碼不得為空",
			},
		)
	}
	token, err := services.Login(inputAccountData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.CommonResponse{
				Err: err.Error(),
			},
		)
	}

	cookie.Set(c, cookie.TokenName, token)

	return c.SendString("Login成功")
}
