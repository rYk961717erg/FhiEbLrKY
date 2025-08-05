// 代码生成时间: 2025-08-05 16:10:52
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    // Import the Fiber package
    "github.com/gofiber/fiber/v2"
)

// AuditLog represents the structure for an audit log entry
type AuditLog struct {
    Timestamp time.Time `json:"timestamp"`
    Method    string    `json:"method"`
    Path      string    `json:"path"`
    IP        string    `json:"ip"`
    Status    int       `json:"status"`
}

// LoggerMiddleware is the middleware function for logging
func LoggerMiddleware(c *fiber.Ctx) error {
    // Start the timer
    start := time.Now()
    defer func() {
        // Calculate the duration
        duration := time.Since(start)
        // Create the audit log entry
        logEntry := AuditLog{
            Timestamp: start,
            Method:    c.Method(),
            Path:      c.Path(),
            IP:        c.IP(),
            Status:    c.Response().StatusCode(),
        }
        // Log the audit log entry
        log.Printf("Audit Log: %+v", logEntry)
    }()
    // Continue to the next middleware
    return c.Next()
}

func main() {
    // Create a new Fiber instance
    app := fiber.New()

    // Use the logger middleware
    app.Use(LoggerMiddleware)

    // Define a route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
