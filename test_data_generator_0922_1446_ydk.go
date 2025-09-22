// 代码生成时间: 2025-09-22 14:46:36
// test_data_generator.go
package main

import (
    "fiber" // 引入fiber框架
    "fmt"
    "math/rand"
    "time"
)

// 初始化随机数生成器
func init() {
    rand.Seed(time.Now().UnixNano())
}

// TestData 定义测试数据的结构
type TestData struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Age       int    `json:"age"`
    IsPremium bool   `json:"is_premium"`
}

// GenerateTestData 生成测试数据
func GenerateTestData() ([]TestData, error) {
    // 定义返回的测试数据切片
    var testData []TestData

    // 生成10条测试数据
    for i := 0; i < 10; i++ {
        // 随机生成用户ID
        userID := fmt.Sprintf("user_%d", i+1)

        // 随机生成用户姓名
        names := []string{"John", "Jane", "Doe", "Smith", "Brown"}
        name := names[rand.Intn(len(names))]

        // 随机生成用户邮箱
        email := fmt.Sprintf("%s@example.com", name)

        // 随机生成用户年龄
        age := rand.Intn(100)

        // 随机判断是否是高级用户
        isPremium := rand.Float32() > 0.5

        // 创建TestData实例并添加到切片
        testData = append(testData, TestData{
            ID:        userID,
            Name:      name,
            Email:     email,
            Age:       age,
            IsPremium: isPremium,
        })
    }
    return testData, nil
}

// setupRoutes 设置路由
func setupRoutes(app *fiber.App) {
    // 测试数据生成器路由
    app.Get("/test-data", func(c *fiber.Ctx) error {
        testData, err := GenerateTestData()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusOK).JSON(testData)
    })
}

func main() {
    // 创建Fiber实例
    app := fiber.New()

    // 设置路由
    setupRoutes(app)

    // 启动服务器
    app.Listen(":3000")
}
