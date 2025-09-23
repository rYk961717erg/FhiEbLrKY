// 代码生成时间: 2025-09-23 11:34:55
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// 定义一个简单的用户模型
type User struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

// newUser 创建一个新的用户
func newUser(id uint, name string) User {
    return User{ID: id, Name: name}
}

// getAllUsers 返回所有用户的列表
func getAllUsers(c *fiber.Ctx) error {
    users := []User{
        newUser(1, "John Doe"),
        newUser(2, "Jane Doe"),
    }
    return c.JSON(users)
}

// getUserByID 根据ID获取用户信息
func getUserByID(c *fiber.Ctx) error {
    id := uint(c.Params("id"))
    // 这里应该添加数据库查询逻辑，此处仅模拟返回
    return c.JSON(newUser(id, "User with ID: " + fmt.Sprint(id)))
}

// main 函数初始化并启动Fiber服务器
func main() {
    app := fiber.New()

    // 定义路由并绑定方法
    app.Get("/users", getAllUsers)
    app.Get("/users/:id", getUserByID)

    // 启动服务器
    app.Listen(":3000")
}
