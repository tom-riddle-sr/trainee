package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Fiber() {
// Basic()
Static()

}


func Basic(){
	// 建立一個fiber實例
	app := fiber.New()

// 設置url處理器,ap響應所有向 "/" 提出的HTTP GET請求
// fiber.Ctx 是一個上下文對象，它包含了有關HTTP請求的信息，並提供了一個用於響應的方法

//匿名函式寫法
// func(parameters) return_type {}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3008")
}
// :value 是一個路由參數，它可以匹配任何在該位置的值
// :name? 是一個可選的路由參數，它可以匹配任何在該位置的值，但是如果沒有提供值，則該路由參數將匹配空字符串
// :api/* 是一個通配符路由參數，它可以匹配任何以 /api/ 開頭的路徑段，並且將匹配的值存儲在通配符路由參數中

func Static(){
	app := fiber.New()
	app.Static("/", "./public")
	app.Listen(":3008")
	fmt.Println("Server is running on port 3008")
}