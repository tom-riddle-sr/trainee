package router

import (
	"trainee/fibertrainee3/handlers"
	"trainee/fibertrainee3/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("insert", middleware.JwtValidate, handlers.StoreData)
	app.Post("login", handlers.Login)
	app.Get("logout", handlers.Logout)

	app.Listen(":3010")
}
