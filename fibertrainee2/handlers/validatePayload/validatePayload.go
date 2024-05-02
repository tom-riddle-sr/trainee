package validatePayload

import (
	"database/sql"
	"fmt"

	"trainee/fibertrainee2/middleware/jwt"
	"trainee/fibertrainee2/model/accountData"
	"trainee/fibertrainee2/service/login"

	"github.com/gofiber/fiber/v2"
)

// step1.🍤insert 連線資料庫&放假資料🍤
// step2.🍤fiber 設置處理器等待響應http請求 🍤
// step3.🍤login 要求登入🍤
// step3.🍤token 成功登入後產生jwt token🍤
// step4. 🍤塞回cookie🍤
// step5. 🍤登出確認jwt token🍤

// struct fiber.Ctx 包含了處理 HTTP 請求所需的所有方法和屬性

func ValidatePayload(DB *sql.DB) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		data := accountData.AccountData{}
		if err := c.BodyParser(&data); err != nil {
			return err
		}

		if data.Account == "" || data.Password == "" {
			// fiber.Map 是一種方便的方式來創建 map[string]interface{}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "帳密不得為空",
			})
		}
		fmt.Println("驗證參數成功")
		login.Login(data.Account, data.Password, DB, c)
		return c.SendString("登入成功")
	})

	app.Use("/logout", func(c *fiber.Ctx) error {
		return jwt.Jwt(c)
	})
	app.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("登出成功")
	})

	app.Listen(":3010")
}
