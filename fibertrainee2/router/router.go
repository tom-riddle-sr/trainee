package router

import (
	"database/sql"
	handlers "trainee/fibertrainee2/handlers/handlers_login"
	"trainee/fibertrainee2/middleware/jwt"

	"github.com/gofiber/fiber/v2"
)

func Router(DB *sql.DB) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", handlers.Login)

	app.Get("/logout", jwt.Jwt(), handlers.Logout)

	app.Listen(":3010")

}
