package handlers

import (
	"fmt"
	"trainee/fibertrainee2/service/login"

	"trainee/fibertrainee2/model"

	"github.com/gofiber/fiber/v2"
)

const (
	jwtCookieName = "token"
)

// 當你在路由處理器中返回一個 error，Fiber 框架會捕獲這個錯誤並將其轉換為一個 HTTP 錯誤響應
// 這樣你就不需要在每個路由處理器中檢查錯誤

func Login(c *fiber.Ctx) error {
	data := model.AccountData{}
	// Go 語言中的所有函數參數都是值傳遞，也就是說，當你傳遞一個變數給一個函數時，這個函數會獲得這個變數的一個副本
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data.Account == "" || data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "帳密不得為空",
		})
	}
	fmt.Println("驗證參數成功")
	tokenString, err := login.Login(data.Account, data.Password)
	if err != nil {
		fmt.Println("登入失敗")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("登入成功")

	cookie := new(fiber.Cookie)
	cookie.Name = jwtCookieName
	cookie.Value = tokenString
	c.Cookie(cookie)
	fmt.Println("存進cookie成功")
	return c.SendString("登入成功")
}

func Logout(c *fiber.Ctx) error {
	token := c.Cookies(jwtCookieName)
	if token == "" {
		return c.Status(400).SendString("未登入")
	}

	c.ClearCookie(jwtCookieName)
	return c.SendString("登出成功")
}

//Cookie &Cookiesㄉ差別?
// c.Cookie：這個方法用來設置一個新的 HTTP cookie
// c.Cookies：這個方法用來讀取一個已存在的 HTTP cookie 的值
