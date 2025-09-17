// 代码生成时间: 2025-09-17 09:33:37
package main

import (
    "fmt"
    "net/url"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// ValidateURL checks if the provided URL is valid
func ValidateURL(c *fiber.Ctx) error {
    // Get the URL from the query parameter
    queryURL := c.Query("url")

    // Check if the URL is empty
    if queryURL == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "URL parameter is missing",
        })
    }

    // Parse the URL to validate its structure
    u, err := url.ParseRequestURI(queryURL)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": fmt.Sprintf("Invalid URL format: %s", err),
        })
    }

    // Check if the scheme is HTTP or HTTPS
    if !strings.HasPrefix(u.Scheme, "http") {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "URL must use HTTP or HTTPS scheme",
        })
    }

    // Return a success message with the URL
    return c.JSON(fiber.Map{
        "message": "URL is valid",
        "url": queryURL,
    })
}

func main() {
    // Create a new Fiber instance
    app := fiber.New()

    // Define the route for URL validation
    app.Get("/validate", ValidateURL)

    // Start the server
    app.Listen(":3000")
}
