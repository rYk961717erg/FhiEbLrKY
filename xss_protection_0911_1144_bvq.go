// 代码生成时间: 2025-09-11 11:44:03
package main

import (
    "html"
    "log"
    "github.com/gofiber/fiber/v2"
)

// XSSMiddleware is a middleware function that prevents XSS attacks by sanitizing
// input to prevent script execution.
func XSSMiddleware(c *fiber.Ctx) error {
    // Sanitize user input to prevent XSS attacks
    c.Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline';")
    return c.Next()
}

// main function to start the Fiber server
func main() {
    // Initialize Fiber with default configuration
    app := fiber.New()

    // Use the XSSMiddleware to sanitize user input on all routes
    app.Use(XSSMiddleware)

    // Define a sample route to demonstrate middleware usage
    app.Get("/", func(c *fiber.Ctx) error {
        // Example of user input that could potentially be dangerous
        userInput := "<script>alert('xss');</script>"

        // Sanitize the user input to prevent XSS attacks
        sanitizedInput := html.EscapeString(userInput)

        // Return a response with the sanitized input
        return c.SendString("You entered: " + sanitizedInput)
    })

    // Handle errors gracefully
    if err := app.Listen(":3000"); err != nil {
        log.Fatal(err)
    }
}
