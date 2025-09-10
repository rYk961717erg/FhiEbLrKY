// 代码生成时间: 2025-09-10 17:07:23
package main

import (
    "math/rand"
    "time"
    "github.com/gofiber/fiber/v2"
)

// RandomNumberGeneratorHandler is a handler function that generates a random number.
// It takes no parameters and returns a JSON response with a random number.
func RandomNumberGeneratorHandler(c *fiber.Ctx) error {
    // Generate a random number between 1 and 100
    randomNumber := rand.Intn(100) + 1
    
    // Return the random number in JSON format
    return c.JSON(fiber.Map{
        "randomNumber": randomNumber,
    })
}

func main() {
    // Create a new Fiber instance
    app := fiber.New()

    // Set the port to listen on
    app.Listen(":3000")

    // Define the route and handler for generating a random number
    app.Get("/random", RandomNumberGeneratorHandler)
}
