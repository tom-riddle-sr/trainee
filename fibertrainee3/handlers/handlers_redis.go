package handlers

import (
	"fmt"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/model/output"
	"trainee/fibertrainee3/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IHandlersRedis interface {
	HandlersSet(c *fiber.Ctx) error
	HandlersGet(c *fiber.Ctx) error
	HandlersDel(c *fiber.Ctx) error
}

type HandlersRedis struct {
	services *services.Services
}

func NewHandlersRedis(services *services.Services) IHandlersRedis {
	return &HandlersRedis{
		services: services,
	}
}

func (h *HandlersRedis) HandlersSet(c *fiber.Ctx) error {
	data := input.SetRedisData{}
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
	if err := h.services.Redis.Set(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("set redis失敗: %w", err).Error(),
		})
	}
	return c.SendString("set redis成功")
}

func (h *HandlersRedis) HandlersGet(c *fiber.Ctx) error {
	data := input.RedisKeyRequest{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}

	if err := h.services.Redis.Get(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("get redis失敗: %w", err).Error(),
		})
	}
	return c.SendString("get redis成功")
}

func (h *HandlersRedis) HandlersDel(c *fiber.Ctx) error {
	data := input.RedisKeyRequest{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(output.ErrResponse{
			Err: fmt.Errorf("解析請求主體失敗: %w", err).Error(),
		})
	}
	if err := h.services.Redis.Del(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
			Err: fmt.Errorf("del redis失敗: %w", err).Error(),
		})
	}
	return c.SendString("del redis成功")
}
