// 代码生成时间: 2025-08-27 02:48:56
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// LoginData 定义登录请求所需的数据结构
type LoginData struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// User 定义用户模型
type User struct {
    Username string
    Password string
}

// MockUser 是一个模拟的用户数据，用于演示
var MockUser = User{
    Username: "admin",
    Password: "password",
}

// validatePassword 用于验证密码是否正确
func validatePassword(providedPassword, storedPassword string) bool {
    return providedPassword == storedPassword
}

// login 处理登录请求
func login(c *fiber.Ctx) error {
    var loginData LoginData
    if err := c.BodyParser(&loginData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": err.Error(),
        })
    }

    // 检查用户名和密码是否匹配
    if loginData.Username == MockUser.Username && validatePassword(loginData.Password, MockUser.Password) {
        return c.JSON(fiber.Map{
            "status":  "success",
            "message": "Login successful",
        })
    } else {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid username or password",
        })
    }
}

func main() {
    app := fiber.New()

    // 设置路由和处理函数
    app.Post("/login", login)

    // 启动服务器
    log.Fatal(app.Listen(":3000"))
}
