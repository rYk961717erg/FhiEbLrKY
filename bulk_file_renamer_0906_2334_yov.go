// 代码生成时间: 2025-09-06 23:34:21
package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "sort"

    "github.com/gofiber/fiber/v2"
)

// 文件重命名工具的配置结构体
type RenameToolConfig struct {
    SourceDir string
    Pattern   string
}

func main() {
    var config RenameToolConfig

    // 解析命令行参数
    flag.StringVar(&config.SourceDir, "dir", "", "Source directory to rename files in")
    flag.StringVar(&config.Pattern, "pattern", "", "Pattern to match and rename files")
    flag.Parse()

    // 检查命令行参数是否有效
    if config.SourceDir == "" || config.Pattern == "" {
        fmt.Println("Both -dir and -pattern options are required")
        return
    }

    // 初始化Fiber
    app := fiber.New()

    // 设置路由处理批量文件重命名
    app.Post("/rename", func(c *fiber.Ctx) error {
        // 绑定请求体到RenameToolConfig结构体
        if err := c.BodyParser(&config); err != nil {
            return err
        }

        // 检查参数
        if config.SourceDir == "" || config.Pattern == "" {
            return c.Status(400).JSON(fiber.Map{
                "error": "Both SourceDir and Pattern are required",
            })
        }

        // 执行文件重命名操作
        err := renameFiles(config.SourceDir, config.Pattern)

        // 返回结果
        if err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.JSON(fiber.Map{
            "message": "Files renamed successfully",
        })
    })

    // 启动服务器
    app.Listen(":3000")
}

// renameFiles 执行文件重命名操作
func renameFiles(sourceDir, pattern string) error {
    // 获取目录中的文件列表
    files, err := os.ReadDir(sourceDir)
    if err != nil {
        return err
    }

    // 按文件名排序
    sort.Slice(files, func(i, j int) bool {
        return files[i].Name() < files[j].Name()
    })

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        // 构建新的文件名
        base := filepath.Base(file.Name())
        if matched, err := filepath.Match(pattern, base); matched && err == nil {
            newName := fmt.Sprintf("new_%s", base) // 简单的重命名逻辑：添加前缀
            newFilePath := filepath.Join(sourceDir, newName)
            oldFilePath := filepath.Join(sourceDir, base)

            // 重命名文件
            if err := os.Rename(oldFilePath, newFilePath); err != nil {
                return err
            }
        }
    }

    return nil
}
