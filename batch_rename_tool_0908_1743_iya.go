// 代码生成时间: 2025-09-08 17:43:31
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// RenameMode 定义重命名模式
type RenameMode struct {
    BaseDir  string
    Pattern  string
    Template string
}

// renameFiles 根据指定的模式和模板批量重命名文件
func renameFiles(mode RenameMode) error {
    files, err := os.ReadDir(mode.BaseDir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        filename := file.Name()
        if strings.Contains(filename, mode.Pattern) {
            newFilename := strings.ReplaceAll(filename, mode.Pattern, mode.Template)
            fromPath := filepath.Join(mode.BaseDir, filename)
            toPath := filepath.Join(mode.BaseDir, newFilename)
            if err := os.Rename(fromPath, toPath); err != nil {
                return fmt.Errorf("failed to rename file %s to %s: %w", filename, newFilename, err)
            }
        }
    }
    return nil
}

// startWebServer 启动Fiber服务器用于重命名文件
func startWebServer() *fiber.App {
    app := fiber.New()

    app.Post("/rename", func(c *fiber.Ctx) error {
        var mode RenameMode
        if err := c.BodyParser(&mode); err != nil {
            return err
        }
        if err := renameFiles(mode); err != nil {
            return err
        }
        return c.SendString("Renaming completed successfully")
    })

    return app
}

func main() {
    app := startWebServer()
    log.Fatal(app.Listen(":3000"))
}
