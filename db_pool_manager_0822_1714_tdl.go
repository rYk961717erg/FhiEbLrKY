// 代码生成时间: 2025-08-22 17:14:53
package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/gofiber/fiber/v2" // Fiber framework
)

// DBConfig holds the database configuration
type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

// DBPool represents the database connection pool
type DBPool struct {
    *sql.DB
    Config DBConfig
}

// NewDBPool creates a new database pool
func NewDBPool(config DBConfig) (*DBPool, error) {
    // Create DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User, config.Password, config.Host, config.Port, config.DBName)

    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Set the maximum number of connections in the idle connection pool.
    db.SetMaxIdleConns(10)

    // Set the maximum number of open connections to the database.
    db.SetMaxOpenConns(100)

    // Ping the database to ensure the connection is alive
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &DBPool{DB: db, Config: config}, nil
}

// Close closes the database connection pool
func (p *DBPool) Close() error {
    return p.DB.Close()
}

func main() {
    // Define database configuration
    config := DBConfig{
        Host:     "localhost",
        Port:     "3306",
        User:     "root",
        Password: "password",
        DBName:   "testdb",
    }

    // Create a new database pool
    dbPool, err := NewDBPool(config)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer dbPool.Close()

    // Create a new Fiber app
    app := fiber.New()

    // Define a route for testing the database connection
    app.Get("/test-db", func(c *fiber.Ctx) error {
        // Perform a query to test the database connection
        var count int
        if err := dbPool.DB.QueryRow("SELECT 1").Scan(&count); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "message": "Database connection is working",
        })
    })

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}
