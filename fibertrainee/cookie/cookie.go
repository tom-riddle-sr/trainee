package cookie

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Cookie() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 從 *fiber.Ctx 中獲取 http.ResponseWriter 和 tokenString
		w := c.Locals("responseWriter").(http.ResponseWriter)
		tokenString := c.Locals("token").(string)

		expire := time.Now().Add(24 * time.Hour)
		cookie := http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expire,
		}
		http.SetCookie(w, &cookie)
		return c.Next()
	}
}
