package login

import (
	"database/sql"
	"fmt"
	"trainee/fibertrainee/common"

	"github.com/gofiber/fiber/v2"
)

func Login(DB *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(&common.Postdata); err != nil {
			return err
		}

		for key, value := range common.Postdata {
			fmt.Sprintf(key, value)
		}

		sqlStatement := fmt.Sprintf(`SELECT * FROM accountdata WHERE account='%s' AND password='%s'`, common.HashString(common.Postdata["account"]), common.HashString(common.Postdata["password"]))
		row := DB.QueryRow(sqlStatement)

		var account, password string
		if err := row.Scan(&account, &password); err != nil {
			fmt.Println("驗證失敗", err)
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}
		fmt.Println("驗證成功")
		return c.Next()
	}
}
