package validatePayload

import (
	"database/sql"
	"fmt"
	"trainee/fibertrainee2/model/accountData"
	"trainee/fibertrainee2/service/login"

	"github.com/gofiber/fiber/v2"
)

func Login(DB *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data := accountData.AccountData{}
		if err := c.BodyParser(&data); err != nil {
			return err
		}

		if data.Account == "" || data.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "帳密不得為空",
			})
		}
		fmt.Println("驗證參數成功")
		login.Login(data.Account, data.Password, DB, c)
		return c.SendString("登入成功")
	}
}

func Logout() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.ClearCookie("token")
		return c.SendString("登出成功")
	}
}
