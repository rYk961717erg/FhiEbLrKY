// 代码生成时间: 2025-09-06 10:49:14
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// FolderOrganizer 结构体，用于封装文件夹整理逻辑
type FolderOrganizer struct {
    BasePath string
}

// NewFolderOrganizer 创建一个 FolderOrganizer 实例
func NewFolderOrganizer(basePath string) *FolderOrganizer {
    return &FolderOrganizer{BasePath: basePath}
}

// Organize 将指定路径下的文件按照扩展名进行分类整理
func (o *FolderOrganizer) Organize() error {
    // 获取文件夹中所有文件和子文件夹
    files, err := os.ReadDir(o.BasePath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        fullPath := filepath.Join(o.BasePath, file.Name())

        // 跳过子文件夹
        if file.IsDir() {
            continue
        }

        // 获取文件扩展名
        extension := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
        if extension == "" {
            continue
        }

        // 创建以扩展名命名的新文件夹
        newDirPath := filepath.Join(o.BasePath, extension)
        if _, err := os.Stat(newDirPath); os.IsNotExist(err) {
            if err := os.MkdirAll(newDirPath, 0755); err != nil {
                return fmt.Errorf("failed to create directory for extension %s: %w", extension, err)
            }
        }

        // 移动文件到新文件夹
        if err := os.Rename(fullPath, filepath.Join(newDirPath, file.Name())); err != nil {
            return fmt.Errorf("failed to move file %s: %w", file.Name(), err)
        }
    }

    return nil
}

func main() {
    app := fiber.New()

    // 设置文件整理的路由
    app.Get("/organize", func(c *fiber.Ctx) error {
        basePath := "./files" // 根据实际情况设置文件夹路径
        organizer := NewFolderOrganizer(basePath)
        if err := organizer.Organize(); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "message": "Folder organized successfully",
        })
    })

    // 启动Fiber服务器
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
