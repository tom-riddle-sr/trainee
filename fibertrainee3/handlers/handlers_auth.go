package handlers

import (
	"fmt"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/model/output"
	"trainee/fibertrainee3/services"
	"trainee/fibertrainee3/tools/cookie"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IHandlersAuth interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type HandlersAuth struct {
	services *services.Services
}

func NewAuth(services *services.Services) IHandlersAuth {
	return &HandlersAuth{
		services: services,
	}
}

func (h *HandlersAuth) Login(c *fiber.Ctx) error {
	inputAccountData := input.CreateAccountData{}
	if err := c.BodyParser(&inputAccountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(inputAccountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			output.ErrResponse{
				Err: fmt.Errorf("驗證輸入資料失敗: %w", err).Error(),
			},
		)
	}

	token, err := h.services.Auth.Login(inputAccountData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			output.ErrResponse{
				Err: err.Error(),
			},
		)
	}

	cookie.Set(c, cookie.TokenName, token)

	return c.SendString("Login成功")
}

func (h *HandlersAuth) Logout(c *fiber.Ctx) error {
	token := cookie.Get(c, cookie.TokenName)
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: "token不存在",
		})
	}
	c.ClearCookie(cookie.TokenName)
	return c.SendString("Logout成功")
}
