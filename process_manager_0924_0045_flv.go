// 代码生成时间: 2025-09-24 00:45:29
package main

import (
    "fmt"
    "os/exec"
    "os"
    "syscall"
    "time"

    "github.com/gofiber/fiber/v2"
)

// ProcessManager 结构体用于存储进程信息
type ProcessManager struct {
    Process *os.Process
}

// NewProcessManager 创建一个新的进程管理器实例
func NewProcessManager(cmd *exec.Cmd) *ProcessManager {
    pm := &ProcessManager{Process: cmd.Process}
    return pm
}

// StartProcess 开始一个进程
func (pm *ProcessManager) StartProcess() error {
    if err := pm.Process.Start(); err != nil {
        return fmt.Errorf("failed to start process: %w", err)
    }
    return nil
}

// StopProcess 停止一个进程
func (pm *ProcessManager) StopProcess() error {
    if err := pm.Process.Signal(syscall.SIGINT); err != nil {
        return fmt.Errorf("failed to stop process: %w", err)
    }
    return nil
}

// WaitProcess 等待一个进程结束
func (pm *ProcessManager) WaitProcess() error {
    if err := pm.Process.Wait(); err != nil {
        return fmt.Errorf("failed to wait for process: %w", err)
    }
    return nil
}

// ProcessHandler 处理进程管理的HTTP请求
func ProcessHandler(c *fiber.Ctx) error {
    cmd := exec.Command("ping", "example.com", "-c", "4")
    pm := NewProcessManager(cmd)

    err := pm.StartProcess()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }

    time.Sleep(2 * time.Second) // 给进程一些时间来执行

    err = pm.StopProcess()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }

    return c.Status(fiber.StatusOK).SendString("Process managed successfully")
}

func main() {
    app := fiber.New()

    app.Get("/process", ProcessHandler)

    app.Listen(":3000")
}
