// 代码生成时间: 2025-09-22 22:03:02
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "html"
)

// XSSProtection Middleware to prevent cross-site scripting attacks
func XSSProtection(c *fiber.Ctx) error {
    // Set HTTP headers to prevent XSS
    c.Set("X-XSS-Protection", "1; mode=block")
    c.Set("Content-Security-Policy", "default-src 'self';")
    c.Set("Referrer-Policy", "no-referrer")
    return c.Next()
}

// sanitizeInput sanitizes user input to prevent XSS attacks
func sanitizeInput(input string) string {
    // Use html.EscapeString to escape HTML special characters
    return html.EscapeString(input)
}

func main() {
    // Create a new Fiber instance
    app := fiber.New()

    // Use the XSS protection middleware for all routes
    app.Use(XSSProtection)

    // Define a route to handle GET requests to "/"
    app.Get("/", func(c *fiber.Ctx) error {
        // Simulate user input
        userInput := "<script>alert('XSS')</script>"

        // Sanitize input to prevent XSS
        sanitizedInput := sanitizeInput(userInput)

        // Return the sanitized input to the client
        return c.SendString(sanitizedInput)
    })

    // Start the Fiber server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
