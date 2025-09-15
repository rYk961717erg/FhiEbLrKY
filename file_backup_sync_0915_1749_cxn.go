// 代码生成时间: 2025-09-15 17:49:32
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "log"
    "sync"
    "time"
    
    "github.com/gofiber/fiber/v2"
)

// FileSyncer 结构体用于保存文件同步所需的信息
type FileSyncer struct {
    SourceDir  string
    TargetDir  string
    LastSync   time.Time
    mu         sync.Mutex
}

// NewFileSyncer 创建一个新的FileSyncer实例
func NewFileSyncer(source, target string) *FileSyncer {
    return &FileSyncer{
        SourceDir: source,
        TargetDir: target,
        LastSync:  time.Now(),
    }
}

// Sync 同步源目录到目标目录
func (fs *FileSyncer) Sync() error {
    fs.mu.Lock()
    defer fs.mu.Unlock()

    // 获取源目录文件列表
    files, err := os.ReadDir(fs.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        srcPath := filepath.Join(fs.SourceDir, file.Name())
        dstPath := filepath.Join(fs.TargetDir, file.Name())

        // 检查目标路径是否存在相同的文件
        if _, err := os.Stat(dstPath); os.IsNotExist(err) {
            // 如果不存在，则从源路径复制到目标路径
            if err := fs.copyFile(srcPath, dstPath); err != nil {
                return fmt.Errorf("failed to copy file %s: %w", file.Name(), err)
            }
        } else if file.ModTime().After(fs.LastSync) {
            // 如果存在但是文件被修改过，则更新文件
            if err := fs.copyFile(srcPath, dstPath); err != nil {
                return fmt.Errorf("failed to update file %s: %w", file.Name(), err)
            }
        }
    }

    fs.LastSync = time.Now()
    return nil
}

// copyFile 复制单个文件从源路径到目标路径
func (fs *FileSyncer) copyFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("failed to open source file %s: %w", src, err)
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("failed to create destination file %s: %w", dst, err)
    }
    defer dstFile.Close()

    if _, err := io.Copy(dstFile, srcFile); err != nil {
        return fmt.Errorf("failed to copy file %s to %s: %w", src, dst, err)
    }
    return nil
}

// BackupHandler 处理文件备份请求
func BackupHandler(c *fiber.Ctx) error {
    source := c.Query("source", "")
    target := c.Query("target", "")
    if source == "" || target == "" {
        return c.Status(fiber.StatusBadRequest).SendString("Source and target directories are required")
    }

    fs := NewFileSyncer(source, target)
    if err := fs.Sync(); err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }

    return c.SendString("Backup and sync completed successfully")
}

func main() {
    app := fiber.New()
    app.Get("/backup", BackupHandler)
    log.Fatal(app.Listen(":3000"))
}