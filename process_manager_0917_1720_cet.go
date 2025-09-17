// 代码生成时间: 2025-09-17 17:20:28
package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
    "time"

    "github.com/gofiber/fiber/v2" // Import the Fiber web framework
)

// ProcessManager 结构体用于管理进程
type ProcessManager struct {
    // 存储管理的进程PID
    ProcessList map[int]*os.Process
}

// NewProcessManager 创建一个新的进程管理器实例
func NewProcessManager() *ProcessManager {
    return &ProcessManager{
        ProcessList: make(map[int]*os.Process),
    }
}

// StartProcess 启动一个新的进程
func (pm *ProcessManager) StartProcess(name string, arg ...string) (*os.Process, error) {
    cmd := exec.Command(name, arg...)
    process, err := cmd.Start()
    if err != nil {
        return nil, err
    }
    pm.ProcessList[process.Pid] = process
    return process, nil
}

// StopProcess 停止指定PID的进程
func (pm *ProcessManager) StopProcess(pid int) error {
    process, exists := pm.ProcessList[pid]
    if !exists {
        return fmt.Errorf("process with PID %d does not exist", pid)
    }
    return process.Signal(syscall.SIGTERM)
}

// ListProcesses 返回当前管理的所有进程
func (pm *ProcessManager) ListProcesses() []*os.Process {
    var processes []*os.Process
    for _, process := range pm.ProcessList {
        processes = append(processes, process)
    }
    return processes
}

// SetupRoutes 设置Fiber的路由
func SetupRoutes(app *fiber.App) {
    // 创建进程管理器实例
    pm := NewProcessManager()

    // POST /start - 启动一个新进程
    app.Post("/start", func(c *fiber.Ctx) error {
        name := c.Query("name")
        arg := c.Query("arg")
        _, err := pm.StartProcess(name, arg)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendStatus(fiber.StatusOK)
    })

    // POST /stop - 停止一个进程
    app.Post="/stop", func(c *fiber.Ctx) error {
        pid := c.Query("pid")
        if pid == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "PID is required",
            })
        }
        err := pm.StopProcess(int(pid))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendStatus(fiber.StatusOK)
    })

    // GET /list - 列出所有进程
    app.Get("/list", func(c *fiber.Ctx) error {
        processes := pm.ListProcesses()
        return c.JSON(processes)
    })
}

func main() {
    // 初始化Fiber应用
    app := fiber.New()

    // 设置路由
    SetupRoutes(app)

    // 启动服务器
    app.Listen(":3000")
}
