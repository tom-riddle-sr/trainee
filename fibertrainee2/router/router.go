package router

import (
	"database/sql"
	"trainee/fibertrainee2/handlers/validatePayload"
	"trainee/fibertrainee2/middleware/jwt"

	"github.com/gofiber/fiber/v2"
)

func Router(DB *sql.DB) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", validatePayload.Login(DB))

	app.Get("/logout", jwt.Jwt(), validatePayload.Logout())

	app.Listen(":3010")

}
