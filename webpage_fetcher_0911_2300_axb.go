// 代码生成时间: 2025-09-11 23:00:43
   Features:
   - Fetches webpage content using HTTP GET request.
   - Handles errors properly.
   - Includes comments and documentation for clarity.
   - Follows Go best practices for maintainability and scalability.
*/

package main

import (
    "fmt"
    "net/http"
    "strings"
    "golang.org/x/net/html"
    "github.com/gofiber/fiber/v2"
)

// FetchWebpageContent retrieves the content of a webpage.
func FetchWebpageContent(url string) (string, error) {
    // Perform an HTTP GET request to the provided URL.
    resp, err := http.Get(url)
    if err != nil {
        // Return error if the request fails.
        return "", fmt.Errorf("failed to fetch webpage: %w", err)
    }
    defer resp.Body.Close()

    // Check if the HTTP request was successful.
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to fetch webpage, status code: %d", resp.StatusCode)
    }

    // Read the content of the webpage.
    var content strings.Builder
    content.Grow(int(resp.ContentLength))
    if _, err := content.ReadFrom(resp.Body); err != nil {
        return "", fmt.Errorf("failed to read webpage content: %w", err)
    }

    // Return the webpage content.
    return content.String(), nil
}

// StartServer initializes and starts the Fiber web server.
func StartServer() {
    app := fiber.New()

    // Define a route to handle GET requests for webpage content.
    app.Get("/fetch", func(c *fiber.Ctx) error {
        // Extract the URL parameter from the query string.
        url := c.Query("url")
        if url == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "URL parameter is missing",
            })
        }

        // Fetch the webpage content and handle any errors.
        content, err := FetchWebpageContent(url)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // Return the webpage content in the response.
        return c.Text(content)
    })

    // Start the server on port 3000.
    if err := app.Listen(":3000"); err != nil {
        // Log and exit if the server fails to start.
        fmt.Printf("failed to start server: %s
", err)
        return
    }
}

func main() {
    // Start the web server.
    StartServer()
}
