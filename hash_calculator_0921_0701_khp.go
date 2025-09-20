// 代码生成时间: 2025-09-21 07:01:41
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fiber"
    "fmt"
    "log"
)

// HashCalculator 结构体用于封装哈希计算功能
type HashCalculator struct {
    // 空结构体，仅用于封装方法
}

// NewHashCalculator 创建一个新的哈希计算工具实例
func NewHashCalculator() *HashCalculator {
    return &HashCalculator{}
}

// CalculateSHA256 计算并返回给定字符串的SHA256哈希值
func (hc *HashCalculator) CalculateSHA256(input string) (string, error) {
    // 创建一个新的SHA256哈希对象
    hash := sha256.New()
    
    // 写入输入数据
    _, err := hash.Write([]byte(input))
    if err != nil {
        return "", err
    }
    
    // 返回哈希值的HEX字符串表示
    return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
    app := fiber.New()
    
    // 创建哈希计算工具实例
    hashCalculator := NewHashCalculator()
    
    // 定义一个路由，用于计算哈希值
    app.Get("/hash", func(c *fiber.Ctx) error {
        // 获取查询参数中的字符串
        inputStr := c.Query("str", "")
        
        if inputStr == "" {
            return c.Status(fiber.StatusBadRequest).SendString("Please provide a string to hash.")
        }
        
        // 计算哈希值
        sha256Str, err := hashCalculator.CalculateSHA256(inputStr)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString("Error calculating hash.")
        }
        
        // 返回哈希值
        return c.JSON(fiber.Map{
            "original": inputStr,
            "hash": sha256Str,
        })
    })
    
    // 启动服务器
    log.Fatal(app.Listen(":3000"))
}
