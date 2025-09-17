// 代码生成时间: 2025-09-18 02:11:55
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/json-iterator/go"
)

// JSON is a global variable to hold the json-iterator instance
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

// Converter is a structure that will handle the JSON conversion
type Converter struct{}

// NewConverter creates a new instance of Converter
func NewConverter() *Converter {
    return &Converter{}
}

// Convert takes a JSON string and returns the string in a modified JSON format
func (c *Converter) Convert(ctx *fiber.Ctx) error {
    // Get the JSON string from the request body
    reqBody := ctx.Context()
    var data map[string]interface{}
    if err := ctx.BodyParser(&data); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":  "Invalid JSON input",
            "message": err.Error(),
        })
    }

    // Convert the data to a JSON string
    jsonString, err := JSON.Marshal(data)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":  "Failed to convert JSON",
            "message": err.Error(),
        })
    }

    // Return the modified JSON string
    return ctx.Status(fiber.StatusOK).Send(string(jsonString))
}

func main() {
    app := fiber.New()

    // Create a new instance of Converter
    converter := NewConverter()

    // Register the /convert endpoint
    app.Post('/convert', converter.Convert)

    // Start the Fiber server
    if err := app.Listen(":3000"); err != nil {
        fmt.Printf("Server failed to start: %s
", err)
    }
}