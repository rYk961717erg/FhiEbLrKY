// 代码生成时间: 2025-08-06 21:13:05
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    \_ "github.com/go-sql-driver/mysql"
    "github.com/gofiber/fiber/v2"
)

// DatabaseConfig is the structure to hold database configuration
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// Database holds the database connection pool
type Database struct {
    *sql.DB
}

// NewDatabase creates a new database connection pool
func NewDatabase(config *DatabaseConfig) (*Database, error) {
    // Construct connection string
    connectionString := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User,
        config.Password,
        config.Host,
        config.Port,
        config.DBName)

    // Open the database connection
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
    }

    // Set maximum lifetime for connections.
    db.SetConnMaxLifetime(5 * time.Minute)
    // Set maximum number of open connections to the database.
    db.SetMaxOpenConns(25)
    // Set maximum number of connections in the idle connection pool.
    db.SetMaxIdleConns(25)

    // Test the connection
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return &Database{db}, nil
}

// Close closes the database connection pool
func (db *Database) Close() error {
    return db.DB.Close()
}

func main() {
    // Define the database configuration
    config := &DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "user",
        Password: "password",
        DBName:   "dbname",
    }

    // Create a new database connection pool
    db, err := NewDatabase(config)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Create a new Fiber instance
    app := fiber.New()

    // Define a route to test the database connection
    app.Get("/testdb", func(c *fiber.Ctx) error {
        // Use the database connection
        err := db.Ping()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Failed to ping database",
                "error":  err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "message": "Database connection established",
        })
    })

    // Start the Fiber server
    address := ":3000"
    log.Printf("Starting server on %s", address)
    if err := app.Listen(address); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
