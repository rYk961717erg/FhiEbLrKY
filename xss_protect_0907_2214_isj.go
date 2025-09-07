// 代码生成时间: 2025-09-07 22:14:22
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/xss"
)

// main 函数启动Fiber应用程序并配置XSS防护中间件
func main() {
    // 创建一个新的Fiber应用程序实例
    app := fiber.New()

    // 使用XSS防护中间件
    // 该中间件将会检查传入的请求并防止XSS攻击
    app.Use(xss.New())

    // 定义一个简单的路由来测试XSS防护
    app.Get("/test", func(c *fiber.Ctx) error {
        // 获取查询参数中的user字段
        user := c.Query("user", "")

        // 响应客户端，显示用户输入，XSS防护中间件将确保输出是安全的
        return c.SendString(fmt.Sprintf("Hello, %s!", user))
    })

    // 启动服务器并打印日志
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: %s", err)
    }
}
