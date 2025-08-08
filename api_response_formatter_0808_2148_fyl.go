// 代码生成时间: 2025-08-08 21:48:54
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// ApiResponse represents the standard response structure for API responses.
type ApiResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents the error details in API responses.
type ErrorInfo struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// NewApiResponse creates a new ApiResponse instance.
func NewApiResponse(data interface{}, err error) *ApiResponse {
    response := &ApiResponse{
        Success: err == nil,
        Data:    data,
    }

    if err != nil {
        response.Error = &ErrorInfo{
            Code:    500, // Internal Server Error by default
            Message: err.Error(),
        }
    }

    return response
}

// main function to run the API server using Fiber.
func main() {
    app := fiber.New()

    // Define an example endpoint to demonstrate ApiResponse usage.
    app.Get("/example", func(c *fiber.Ctx) error {
        // Simulate some data retrieval.
        data := "Sample data"

        // Create a new ApiResponse with the data.
        response := NewApiResponse(data, nil)

        // Return the ApiResponse as JSON.
        return c.JSON(response)
    })

    // Define an example error endpoint.
    app.Get("/error", func(c *fiber.Ctx) error {
        // Simulate an error scenario.
        return NewApiResponse(nil, fmt.Errorf("an error occurred"))
    })

    // Start the Fiber server on port 3000.
    log.Println("Server is running on :3000")
    if err := app.Listen(":3000"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}