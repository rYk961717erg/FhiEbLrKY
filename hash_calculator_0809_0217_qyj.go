// 代码生成时间: 2025-08-09 02:17:20
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fiber/fiber/v2"
    "net/http"
    "os"
)

// HashCalculator 结构体用于处理哈希计算
type HashCalculator struct {
    app *fiber.App
}

// NewHashCalculator 创建一个新的 HashCalculator 实例
func NewHashCalculator() *HashCalculator {
    app := fiber.New()
    return &HashCalculator{
        app: app,
    }
}

// CalculateHash 提供 HTTP 方法来计算给定输入的 SHA-256 哈希值
func (h *HashCalculator) CalculateHash() error {
    h.app.Get("/hash", func(c *fiber.Ctx) error {
        input := c.Query("input")
        if input == "" {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": "Input parameter 'input' is required.",
            })
        }

        // 计算 SHA-256 哈希值
        hash := sha256.Sum256([]byte(input))
        hexHash := hex.EncodeToString(hash[:])

        // 响应哈希值
        return c.JSON(fiber.Map{
            "input": input,
            "hash": hexHash,
        })
    })
    return nil
}

// Start 启动 HTTP 服务
func (h *HashCalculator) Start(port string) error {
    return h.app.Listen(port)
}

func main() {
    // 创建哈希计算工具实例
    calculator := NewHashCalculator()
    // 添加路由和中间件
    if err := calculator.CalculateHash(); err != nil {
        fmt.Printf("Error setting up routes: %s
", err)
        os.Exit(1)
    }
    // 启动服务
    if err := calculator.Start(":3000"); err != nil {
        fmt.Printf("Error starting server: %s
", err)
        os.Exit(1)
    }
}
