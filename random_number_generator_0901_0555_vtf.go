// 代码生成时间: 2025-09-01 05:55:15
package main

import (
    "crypto/rand"
    "encoding/binary"
    "fmt"
    "math/big"
    
    "github.com/gofiber/fiber/v2"
)

// RandomNumberGeneratorHandler is a fiber handler function that generates a random number
// and returns it as a JSON response.
func RandomNumberGeneratorHandler(c *fiber.Ctx) error {
    // Generate a random number between 1 and 100
    randomNumber, err := GenerateRandomNumber(1, 100)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to generate random number.",
        })
    }

    // Return the random number as a JSON response
    return c.JSON(fiber.Map{
        "random_number": randomNumber,
    })
}

// GenerateRandomNumber generates a random number between min and max
func GenerateRandomNumber(min, max int) (int, error) {
    // Ensure max is greater than min
    if max <= min {
        return 0, fmt.Errorf("max must be greater than min")
    }

    // Generate a random number
    randNum, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
    if err != nil {
        return 0, err
    }

    // Add min to the generated random number and return
    return int(randNum.Int64()) + min, nil
}

func main() {
    app := fiber.New()

    // Register the random number generator handler
    app.Get("/random-number", RandomNumberGeneratorHandler)

    // Start the Fiber server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
