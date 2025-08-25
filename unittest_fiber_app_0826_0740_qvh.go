// 代码生成时间: 2025-08-26 07:40:15
package main

import (
    "fmt"
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/utils"
    "github.com/stretchr/testify/assert"
)

// App is a struct to hold Fiber instance.
type App struct {
    Fiber *fiber.App
}

// NewApp creates a new instance of App.
func NewApp() *App {
    return &App{
        Fiber: fiber.New(),
    }
}

// SetupTest initializes the test environment.
func SetupTest() *App {
    app := NewApp()
    return app
}

// TestGetRoute tests the route setup by the Fiber app.
func TestGetRoute(t *testing.T) {
    app := SetupTest()
    app.Fiber.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    assert := assert.New(t)
    res, err := app.Fiber.Test("GET", "/test")
    assert.NoError(err)
    assert.Equal(fiber.StatusOK, res.StatusCode)
    assert.Equal("Hello, World!", res.Body.String())
}

// TestMain runs the Fiber app if the test is run directly.
func TestMain(m *testing.M) {
    app := SetupTest()
    app.Fiber.Listen(":8080\)
    fmt.Println("Server is running...")
    m.Run()
}
