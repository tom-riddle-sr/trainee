package cookie

import "github.com/gofiber/fiber/v2"

const TokenName = "loginData"

func Set(c *fiber.Ctx, key, value string) {
	cookie := new(fiber.Cookie)
	cookie.Name = key
	cookie.Value = value
	c.Cookie(cookie)
}

func Get(c *fiber.Ctx, key string) string {
	return c.Cookies(key)
}
