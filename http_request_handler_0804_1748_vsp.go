// 代码生成时间: 2025-08-04 17:48:21
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// App holds Fiber app instance
type App struct {
    *fiber.App
}

// NewApp creates a new instance of App
func NewApp() *App {
    return &App{
        App: fiber.New(fiber.Config{}),
    }
}

// Start starts the Fiber app
func (a *App) Start() {
    if err := a.App.Listen(":3000"); err != nil && err != fiber.ErrServerClosed {
        log.Fatalf("An error occurred while starting the server: "%s"", err)
    }
}

// PingHandler responds with 'pong'
func PingHandler(c *fiber.Ctx) error {
    return c.SendString("pong")
}

// main function starts the Fiber app
func main() {
    app := NewApp()

    // Register handlers
    app.Get("/ping", PingHandler)

    // Start the Fiber app
    app.Start()
}
