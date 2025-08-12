// 代码生成时间: 2025-08-12 19:19:20
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "strings"
    "html"
)

// sanitizeInput sanitizes input to prevent XSS attacks
func sanitizeInput(input string) string {
    // Convert the input string to a Reader
    reader := strings.NewReader(input)
    // Create a new StringWriter to hold the sanitized string
    writer := strings.Builder{}
    // Define a list of allowed tags
    allowedTags := []string{"b", "i", "u", "strong", "em", "a", "p", "br", "ul", "ol", "li"}
    // Sanitize the input by only allowing the allowed tags
    html.EscapeString(reader, writer, allowedTags)
    return writer.String()
}

// indexHandler handles the index route
func indexHandler(c *fiber.Ctx) error {
    // Get the user input from the query parameter
    userInput := c.Query("input", "")
    
    // Sanitize the user input
    sanitizedInput := sanitizeInput(userInput)
    
    // Render the sanitized input to the user
    return c.SendString(fmt.Sprintf("<h1>Sanitized Input:</h1><p>%s</p>", sanitizedInput))
}

func main() {
    // Create a new Fiber app
    app := fiber.New()
    
    // Define the index route
    app.Get("/", indexHandler)
    
    // Start the Fiber server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
        return
    }
}
