// 代码生成时间: 2025-09-14 00:33:39
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "os/exec"
    "strings"
)

// SystemMonitorHandler 定义系统性能监控的路由处理器
func SystemMonitorHandler(c *fiber.Ctx) error {
    // 获取CPU和内存使用率
    cpu, mem, err := GetSystemStats()
    if err != nil {
        // 错误处理
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to get system stats",
        })
    }
    
    return c.JSON(fiber.Map{
        "cpu_usage": cpu,
        "memory_usage": mem,
    })
}

// GetSystemStats 获取系统的CPU和内存使用率
func GetSystemStats() (float64, float64, error) {
    // 使用命令获取CPU和内存使用率
    cpuCmd := "top -bn1 | grep 'Cpu(s)' | sed 's/.*, *\([0-9.]*\)%* id.*/\1/' | awk '{print 100 - $1}'"
    memCmd := "free -m | awk 'NR==2{printf "%.2f", $3/$2 * 100.0}'"
    
    // 执行命令
    cpuOut, err := exec.Command("/bin/sh", "-c", cpuCmd).Output()
    if err != nil {
        return 0, 0, err
    }
    memOut, err := exec.Command("/bin/sh", "-c", memCmd).Output()
    if err != nil {
        return 0, 0, err
    }
    
    // 将输出转换为浮点数
    cpuUsage, err := strconv.ParseFloat(strings.TrimSpace(string(cpuOut)), 64)
    if err != nil {
        return 0, 0, err
    }
    memUsage, err := strconv.ParseFloat(strings.TrimSpace(string(memOut)), 64)
    if err != nil {
        return 0, 0, err
    }
    
    return cpuUsage, memUsage, nil
}

func main() {
    app := fiber.New()
    
    // 定义路由
    app.Get("/system/monitor", SystemMonitorHandler)
    
    // 启动服务器
    fmt.Println("Server is running on :3000")
    if err := app.Listen(":3000"); err != nil {
        fmt.Println(err)
    }
}