// 代码生成时间: 2025-08-29 03:03:50
package main

import (
    "fmt"
    "math"
    "strconv"

    "github.com/gofiber/fiber/v2"
)

// MathCalculator contains methods for various mathematical operations.
type MathCalculator struct{}

// Add handles the addition operation.
func (mc *MathCalculator) Add(c *fiber.Ctx) error {
    a, err := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, err := strconv.ParseFloat(c.Query("b", "0"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input parameters."
        })
    }
    result := a + b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// Subtract handles the subtraction operation.
func (mc *MathCalculator) Subtract(c *fiber.Ctx) error {
    a, err := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, err := strconv.ParseFloat(c.Query("b", "0"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input parameters."
        })
    }
    result := a - b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// Multiply handles the multiplication operation.
func (mc *MathCalculator) Multiply(c *fiber.Ctx) error {
    a, err := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, err := strconv.ParseFloat(c.Query("b", "0"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input parameters."
        })
    }
    result := a * b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// Divide handles the division operation with error handling for division by zero.
func (mc *MathCalculator) Divide(c *fiber.Ctx) error {
    a, err := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, err := strconv.ParseFloat(c.Query("b", "1"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input parameters."
        })
    }
    if b == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot divide by zero."
        })
    }
    result := a / b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// Main function to setup and start the Fiber server.
func main() {
    app := fiber.New()
    
    calc := MathCalculator{}
    
    // Define routes for each mathematical operation.
    app.Get("/add", calc.Add)
    app.Get("/subtract", calc.Subtract)
    app.Get("/multiply", calc.Multiply)
    app.Get("/divide", calc.Divide)
    
    // Start the server.
    if err := app.Listen(":8080"); err != nil {
        fmt.Println("Error starting server: &{err}")
    }
}