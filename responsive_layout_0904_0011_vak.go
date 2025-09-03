// 代码生成时间: 2025-09-04 00:11:21
package main

import (
    "fiber/fiber"
    "fmt"
    "net/http"
)

// ResponseHandler is a function signature for handling HTTP responses.
type ResponseHandler func(*fiber.Ctx) error

// appHandler handles requests for the application.
// It returns HTML content with different layouts based on the request's User-Agent.
func appHandler(c *fiber.Ctx) error {
    useragent := c.Get("User-Agent")
    // Check if the user agent is a mobile device
    isMobile := contains(useragent, "Mobi") || contains(useragent, "Android") || contains(useragent, "iPhone")

    // Based on the device type, return the appropriate layout
    if isMobile {
        return c.SendString("Mobile layout")
    } else {
        return c.SendString("Desktop layout")
    }
}

// contains checks if a substring exists within the main string.
func contains(s, substr string) bool {
    return strings.Contains(s, substr)
}

func main() {
    // Create a new Fiber instance.
    app := fiber.New()

    // Define a route for the application that handles GET requests.
    app.Get("/", appHandler)

    // Start the Fiber server.
    // The server will listen on port 3000.
    if err := app.Listen(":3000"); err != nil {
        fmt.Printf("Error starting server: %s
", err)
        return
    }
}
