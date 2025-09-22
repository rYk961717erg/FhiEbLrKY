// 代码生成时间: 2025-09-22 10:36:29
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/gofiber/fiber/v2" // Fiber framework
)

// DBConfig holds the configuration for the database connection.
type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Database string
}

// DBPool represents a database connection pool.
type DBPool struct {
    *sql.DB
}

// NewDBPool creates a new database pool with the given configuration.
func NewDBPool(config DBConfig) (*DBPool, error) {
    // Create DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User, config.Password, config.Host, config.Port, config.Database)

    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Set the maximum number of connections in the idle connection pool.
    db.SetMaxIdleConns(10)

    // Set the maximum number of open connections to the database.
    db.SetMaxOpenConns(100)

    // Set the connection maximum life time.
    db.SetConnMaxLifetime(5 * time.Minute)

    // Test the database connection.
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &DBPool{db}, nil
}

func main() {
    // Define the database configuration.
    config := DBConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "your_username",
        Password: "your_password",
        Database: "your_database",
    }

    // Create a new database pool.
    dbPool, err := NewDBPool(config)
    if err != nil {
        log.Fatalf("Failed to create database pool: %v", err)
    }
    defer dbPool.Close() // Ensure the pool is closed when the app shuts down.

    // Initialize the Fiber app.
    app := fiber.New()

    // Define a route to demonstrate the use of the database pool.
    app.Get("/", func(c *fiber.Ctx) error {
        result, err := dbPool.Query("SELECT 1")
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to query the database",
            })
        }
        defer result.Close()

        return c.SendString("Database connection established successfully.")
    })

    // Start the Fiber server.
    log.Fatal(app.Listen(":3000"))
}