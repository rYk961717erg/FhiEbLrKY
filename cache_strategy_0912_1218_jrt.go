// 代码生成时间: 2025-09-12 12:18:35
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cache"
    "time"
)

// CacheStrategyMiddleware 实现缓存策略的中间件
func CacheStrategyMiddleware(app *fiber.App) *fiber.App {
    // 使用Fiber中间件缓存静态文件，例如图片、CSS、JavaScript文件等
    // 有效期设置为1天
    app.Use(cache.New(cache.Config{
        Expiration: 24 * time.Hour,
        CacheControl: true,
    }))

    return app
}

// main 函数是程序的入口点
func main() {
    app := fiber.New()

    // 应用缓存策略中间件
    CacheStrategyMiddleware(app)

    // 定义路由和处理函数
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to the Fiber Cache Strategy Example!")
    })

    // 启动服务器
    fmt.Println("Server started on :3000")
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
