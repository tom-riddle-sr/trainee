package router

import (
	"fmt"
	"trainee/fibertrainee3/handlers"
	"trainee/fibertrainee3/middleware"

	"trainee/fibertrainee3/model/output"

	"github.com/gofiber/fiber/v2"
)

func Router(handlers *handlers.Handlers) {
	app := fiber.New(fiber.Config{
		ErrorHandler: nil,
	})
	// 設置 panic handler
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			// panic 發生 return panic error response
			if r := recover(); r != nil {
				// 嘗試將 panic 的值轉換為一個 error
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				c.Status(fiber.StatusInternalServerError).JSON(output.ErrResponse{
					Err: fmt.Errorf("panic錯誤: %w", err).Error(),
				})
			}
		}()
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("這是test panic")
	})

	api1 := app.Group("/auth")
	api1.Post("/login", handlers.Auth.Login)
	api1.Get("/logout", handlers.Auth.Logout)

	api := app.Group("/account", middleware.JwtValidate)
	api.Post("create", handlers.Account.CreateAccount)
	api.Post("update", handlers.Account.UpdateAccount)
	api.Post("delete", handlers.Account.DeleteAccount)

	app.Listen(":3010")
}
