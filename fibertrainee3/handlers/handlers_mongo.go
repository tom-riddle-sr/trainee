package handlers

import (
	"fmt"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/model/output"
	"trainee/fibertrainee3/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IHandlersMongo interface {
	HandlersInsert(c *fiber.Ctx) error
	HandlersUpdate(c *fiber.Ctx) error
	HandlersFind(c *fiber.Ctx) error
	HandlersDelete(c *fiber.Ctx) error
}
type HandlersMongo struct {
	services *services.Services
}

func NewHandlersMongo(services *services.Services) IHandlersMongo {
	return &HandlersMongo{
		services: services,
	}
}

func (h *HandlersMongo) HandlersInsert(c *fiber.Ctx) error {
	data := input.InsertMongoData{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			output.ErrResponse{
				Err: fmt.Errorf("驗證輸入資料失敗: %w", err).Error(),
			},
		)
	}

	if err := h.services.Mongo.Insert(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("insert失敗: %w", err).Error(),
		})
	}
	return c.SendString("insert成功")
}

func (h *HandlersMongo) HandlersUpdate(c *fiber.Ctx) error {
	data := input.UpdateMongoData{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	if err := h.services.Mongo.Update(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("update失敗: %w", err).Error(),
		})
	}
	return c.SendString("update成功")
}

func (h *HandlersMongo) HandlersFind(c *fiber.Ctx) error {
	data := input.MongoIDRequest{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	if err := h.services.Mongo.FindOne(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("find失敗: %w", err).Error(),
		})
	}
	return c.SendString("find成功")
}

func (h *HandlersMongo) HandlersDelete(c *fiber.Ctx) error {
	data := input.MongoIDRequest{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	if err := h.services.Mongo.DeleteOne(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("delete失敗: %w", err).Error(),
		})
	}
	return c.SendString("delete成功")
}
