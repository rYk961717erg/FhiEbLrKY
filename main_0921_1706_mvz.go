// 代码生成时间: 2025-09-21 17:06:46
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

// DBClient is the global variable for the database client
var DBClient *gorm.DB

// SetupDatabase initializes the database connection
func SetupDatabase() error {
    // Connect to SQLite database
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    // Migrate the schema
    if err := db.AutoMigrate(&User{}); err != nil {
        return err
    }

    // Set the global variable
    DBClient = db
    return nil
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
    var newUser User
    if err := c.BodyParser(&newUser); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Validate user data
    if newUser.Name == "" || newUser.Email == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Name and Email are required",
        })
    }

    // Create user
    if err := DBClient.Create(&newUser).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(newUser)
}

func main() {
    // Set up the database
    if err := SetupDatabase(); err != nil {
        fmt.Printf("Failed to set up database: %s
", err)
        return
    }

    // Create a new Fiber app
    app := fiber.New()

    // Set up routes
    app.Post("/users", CreateUser)

    // Start the server
    app.Listen(":3000")
}
