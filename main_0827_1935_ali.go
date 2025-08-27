// 代码生成时间: 2025-08-27 19:35:53
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User represents the data model for a user
type User struct {
    gorm.Model
    Name  string `json:"name"`
    Email string `json:"email"`
}

// DB is a global instance of the GORM database connection
var DB *gorm.DB

func main() {
    // Setup the SQLite database connection
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    DB.AutoMigrate(&User{})

    // Create a new Fiber instance
    app := fiber.New()

    // Define the route for creating a new user
    app.Post("/users", CreateUser)

    // Start the server
    app.Listen(":3000")
}

// CreateUser handles the HTTP POST request for creating a new user
func CreateUser(c *fiber.Ctx) error {
    // Bind and validate the request body
    user := new(User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": err.Error(),
        })
    }

    // Save the new user to the database
    if err := DB.Create(user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": err.Error(),
        })
    }

    // Return the created user as JSON
    return c.Status(fiber.StatusOK).JSON(user)
}