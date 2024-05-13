package handlers

import (
	"fmt"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/model/output"
	"trainee/fibertrainee3/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IHandlersAccount interface {
	CreateAccount(c *fiber.Ctx) error
	UpdateAccount(c *fiber.Ctx) error
	DeleteAccount(c *fiber.Ctx) error
}

type HandlersAccount struct {
	services *services.Services
}

func NewAccount(services *services.Services) IHandlersAccount {
	return &HandlersAccount{
		services: services,
	}
}

func (h *HandlersAccount) CreateAccount(c *fiber.Ctx) error {
	accountData := input.CreateAccountData{}
	if err := c.BodyParser(&accountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(accountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			output.ErrResponse{
				Err: fmt.Errorf("驗證輸入資料失敗: %w", err).Error(),
			},
		)
	}

	if err := h.services.Account.StoreData(&accountData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("create失敗: %w", err).Error(),
		})
	}
	return c.SendString("create成功")
}

func (h *HandlersAccount) UpdateAccount(c *fiber.Ctx) error {
	accountData := input.UpdateAccountData{}
	if err := c.BodyParser(&accountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(accountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			output.ErrResponse{
				Err: fmt.Errorf("驗證輸入資料失敗: %w", err).Error(),
			},
		)
	}

	if err := h.services.Account.UpdateData(accountData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("update失敗: %w", err).Error(),
		})
	}
	return c.SendString("update成功")
}

func (h *HandlersAccount) DeleteAccount(c *fiber.Ctx) error {
	accountData := input.IDRequest{}
	if err := c.BodyParser(&accountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(accountData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			output.ErrResponse{
				Err: fmt.Errorf("驗證輸入資料失敗: %w", err).Error(),
			},
		)
	}

	if err := h.services.Account.DeleteData(accountData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("delete: %w", err).Error(),
		})
	}
	return c.SendString("delete成功")
}
