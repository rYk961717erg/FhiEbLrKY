// 代码生成时间: 2025-08-15 13:52:29
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// CheckAccess defines a function type for access control checks
type CheckAccess func(c *fiber.Ctx) bool

// AccessControlMiddleware is a middleware that checks if the user has access
func AccessControlMiddleware(check CheckAccess) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if !check(c) {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "error": "Access denied",
            })
        }
        return c.Next()
    }
}

// ExampleCheckAccess is a simple access check function example
func ExampleCheckAccess(c *fiber.Ctx) bool {
    // Example access check logic, replace with real logic
    return c.Get("Authorization") == "secret-token"
}

func main() {
    app := fiber.New()

    // Register the access control middleware with your access check function
    app.Use(AccessControlMiddleware(ExampleCheckAccess))

    // Define a protected route
    app.Get("/protected", func(c *fiber.Ctx) error {
        return c.SendString("This is a protected resource")
    })

    // Start the server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
