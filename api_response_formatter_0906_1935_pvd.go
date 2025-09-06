// 代码生成时间: 2025-09-06 19:35:46
Features:
- Code structure is clear and easy to understand.
- Includes appropriate error handling.
- Contains necessary comments and documentation.
- Follows Go best practices.
- Ensures code maintainability and scalability.
*/

package main

import (
    "fmt"
    "log"

    // Import the Fiber framework
    "github.com/gofiber/fiber/v2"
)

// ErrorResponse represents the structure for error responses.
type ErrorResponse struct {
    Success bool        "json:"success" example:"false""
    Message string     "json:"message" example:"An error occurred."
    Data    interface{} "json:"data" example:""
}

// SuccessResponse represents the structure for success responses.
type SuccessResponse struct {
    Success bool        "json:"success" example:"true""
    Message string     "json:"message" example:"Operation successful."
    Data    interface{} "json:"data" example:""
}

// Formatter contains methods to format API responses.
type Formatter struct {}

// NewFormatter creates a new instance of the Formatter.
func NewFormatter() *Formatter {
    return &Formatter{}
}

// FormatSuccess formats a success response.
func (f *Formatter) FormatSuccess(c *fiber.Ctx, data interface{}, message string) error {
    resp := SuccessResponse{
        Success: true,
        Message: message,
        Data:    data,
    }
    return c.Status(fiber.StatusOK).JSON(resp)
}

// FormatError formats an error response.
func (f *Formatter) FormatError(c *fiber.Ctx, message string, data interface{}) error {
    resp := ErrorResponse{
        Success: false,
        Message: message,
        Data:    data,
    }
    return c.Status(fiber.StatusInternalServerError).JSON(resp)
}

func main() {
    // Create a new Fiber app
    app := fiber.New()

    // Create a new formatter instance
    formatter := NewFormatter()

    // Define a route to demonstrate the response formatter
    app.Get("/format-response", func(c *fiber.Ctx) error {
        // Simulate a successful response
        if err := formatter.FormatSuccess(c, "some data", "This is a success message."); err != nil {
            return err
        }
        return nil
    })

    // Define a route to demonstrate error handling
    app.Get("/error-response", func(c *fiber.Ctx) error {
        // Simulate an error response
        return formatter.FormatError(c, "An error occurred.", nil)
    })

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}
