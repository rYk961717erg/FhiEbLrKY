// 代码生成时间: 2025-09-03 09:18:24
package main

import (
	"fmt"
	"log"
	"net/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

// User represents a user entity in the database.
type User struct {
	gorm.Model
	Name    string
	Email   string `gorm:"type:varchar(100);uniqueIndex"`
	// Add more fields as needed
}

// Database is a global variable to hold the GORM DB connection.
var Database *gorm.DB

// SetupDatabase initializes the database connection and migrates the schema.
func SetupDatabase() {
	var err error
	Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	n := Database.AutoMigrate(&User{})
	if n == nil {
		log.Println("Could not migrate schema: ", err)
	}
}

// GetUserHandler is the handler for the GET /user/{id} endpoint.
func GetUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	// Use the Find method to retrieve a single user by ID, which prevents SQL injection.
	user := User{}
	result := Database.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(user)
}

func main() {
	SetupDatabase()

defender := fiber.New()

defender.Get("/user/:id", GetUserHandler)

	log.Fatal(defender.Listen(":3000"))
}
