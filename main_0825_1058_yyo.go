// 代码生成时间: 2025-08-25 10:58:18
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

// User represents a user data model
type User struct {
    ID     uint   "json:"id" xml:"id""
    Name   string "json:"name" xml:"name""
    Email  string "json:"email" xml:"email""
    Active bool   "json:"active" xml:"active""
}

// NewUser represents a user data model for creation
type NewUser struct {
    Name  string `json:"name" xml:"name"`
    Email string `json:"email" xml:"email"`
    Active bool   `json:"active" xml:"active"`
}

// Routes defines the routes for the application
func Routes(app *fiber.App) {
    // Enable CORS
    app.Use(cors.New())

    // Users endpoint
    app.Get("/users", GetAllUsers)
    app.Post("/users", CreateUser)
    app.Get("/users/:id", GetUserByID)
    app.Patch("/users/:id", UpdateUser)
    app.Delete("/users/:id", DeleteUser)
}

func main() {
    // Create a new Fiber app
    app := fiber.New()

    // Set up routes
    Routes(app)

    // Start the server
    app.Listen(":3000")
}

// GetAllUsers responds with a list of all users
func GetAllUsers(c *fiber.Ctx) error {
    // TODO: Fetch users from the database
    return c.JSON(fiber.Map{
        "message": "List of users",
        "data": []User{
            {ID: 1, Name: "John Doe", Email: "john@example.com", Active: true},
            {ID: 2, Name: "Jane Doe", Email: "jane@example.com", Active: false},
        },
    })
}

// CreateUser adds a new user to the database
func CreateUser(c *fiber.Ctx) error {
    // Extract JSON body
    var newUser NewUser
    if err := c.BodyParser(&newUser); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // TODO: Save the new user to the database

    return c.JSON(fiber.Map{
        "message": "User created",
        "data": newUser,
    })
}

// GetUserByID retrieves a user by their ID
func GetUserByID(c *fiber.Ctx) error {
    // Extract user ID from the route param
    userID := c.Params("id")

    // TODO: Fetch the user from the database

    return c.JSON(fiber.Map{
        "message": "User retrieved",
        "data": User{ID: 1, Name: "John Doe", Email: "john@example.com", Active: true},
    })
}

// UpdateUser updates a user's information
func UpdateUser(c *fiber.Ctx) error {
    // Extract user ID from the route param
    userID := c.Params("id")

    // Extract JSON body
    var updateUser User
    if err := c.BodyParser(&updateUser); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // TODO: Update the user in the database

    return c.JSON(fiber.Map{
        "message": "User updated",
        "data": updateUser,
    })
}

// DeleteUser removes a user from the database
func DeleteUser(c *fiber.Ctx) error {
    // Extract user ID from the route param
    userID := c.Params("id")

    // TODO: Delete the user from the database

    return c.SendStatus(fiber.StatusNoContent)
}
