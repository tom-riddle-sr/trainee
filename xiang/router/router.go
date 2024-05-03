package router

import (
	"trainee/fibertrainee2/handlers"
	"trainee/fibertrainee2/middleware/jwt"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", handlers.Login)

	app.Get("/logout", jwt.New(), handlers.Logout)

	app.Listen(":3010")
}
