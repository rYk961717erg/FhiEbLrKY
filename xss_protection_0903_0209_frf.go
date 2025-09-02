// 代码生成时间: 2025-09-03 02:09:02
It uses `bluemonday.UGCPolicy()` to sanitize user input on a POST endpoint.
*/

package main

import (
    "fmt"
    "net/http"
    "strings"
    "github.com/gofiber/fiber"
    "github.com/microcosm-cc/bluemonday"
)

// sanitizeInput sanitizes user input to prevent XSS attacks.
func sanitizeInput(input string) string {
    // Create a new policy instance
    policy := bluemonday.UGCPolicy()
    // Sanitize the input
    return policy.Sanitize(input)
}

// main function to set up the Fiber server and routes.
func main() {
    app := fiber.New()

    // Define a GET route for demonstration
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("This is a simple XSS protection service.")
    })

    // Define a POST route to handle user input and sanitize it.
    app.Post("/sanitize", func(c *fiber.Ctx) error {
        // Get the user input from the request
        input := c.FormValue("userInput")

        // Sanitize the input to prevent XSS
        sanitizedInput := sanitizeInput(input)

        // Respond with the sanitized input
        return c.JSON(fiber.Map{
            "originalInput": input,
            "sanitizedInput": sanitizedInput,
        })
    }, fiber.ErrorHandlerfunc(err error) error {
        // Custom error handling
        return c.Status(http.StatusInternalServerError).SendString("There was an error: " + err.Error())
    })

    // Start the Fiber server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
        return
    }
}
