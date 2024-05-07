package handlers

import (
	"trainee/fibertrainee3/model"
	"trainee/fibertrainee3/repository"

	"trainee/fibertrainee3/database/mysql"

	"github.com/gofiber/fiber/v2"
)

func StoreData(c *fiber.Ctx) error {
	if err := repository.Insert(mysql.GetDB()); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.CommonResponse{
			Err: err.Error(),
		})
	}
	return c.SendString("Insert成功")
}
