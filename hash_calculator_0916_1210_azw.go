// 代码生成时间: 2025-09-16 12:10:41
package main

import (
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "log"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// CalculateSHA1 calculates the SHA1 hash of the given input string.
func CalculateSHA1(input string) (string, error) {
    hash := sha1.New()
    _, err := hash.Write([]byte(input))
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
    app := fiber.New()

    // Define a route to calculate SHA1 hash for the input string.
    app.Post("/hash", func(c *fiber.Ctx) error {
        // Extract the input string from the request body.
        input := c.Query("input")
        if input == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Input string is required",
            })
        }

        // Calculate the SHA1 hash.
        sha1Hash, err := CalculateSHA1(input)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": fmt.Sprintf("Failed to calculate hash: %s", err),
            })
        }

        // Return the calculated SHA1 hash.
        return c.JSON(fiber.Map{
            "input": input,
            "hash": sha1Hash,
        })
    })

    // Start the Fiber server.
    log.Fatal(app.Listen(":3000"))
}
