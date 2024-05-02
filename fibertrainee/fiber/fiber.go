package fiber

import (
	"database/sql"
	"trainee/fibertrainee/common"
	"trainee/fibertrainee/cookie"
	"trainee/fibertrainee/login"
	"trainee/fibertrainee/token"

	"github.com/gofiber/fiber/v2"
)

// struct fiber.Ctx 包含了處理 HTTP 請求所需的所有方法和屬性

func Fiber(DB *sql.DB) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		loginErr := login.Login(DB)(c)
		if loginErr != nil {
			return loginErr
		}

		tokenErr := token.Token()(c)
		if tokenErr != nil {
			return tokenErr
		}

		c.Cookies("token", common.Jwttoken)

		cookieErr := cookie.Cookie()(c)
		if cookieErr != nil {
			return cookieErr
		}

		return c.SendString("yeah i got u data")
	})

	app.Listen(":3010")
}
