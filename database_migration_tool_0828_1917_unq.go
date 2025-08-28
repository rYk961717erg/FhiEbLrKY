// 代码生成时间: 2025-08-28 19:17:54
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "golang.org/x/exp/slices"
    "github.com/go-sql-driver/mysql"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/jmoiron/sqlx"
    "github.com/pressly/goose/v3"
)

// DatabaseConfig holds the database configuration
type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}

// Migration contains the necessary information for a database migration
type Migration struct {
    Version    string
    Direction  goose.Direction
    FileName   string
    AppliedAt int64
}

func main() {
    // Set up database
    dbConfig := DatabaseConfig{
        Host:     "localhost",
        Port:     "3306",
        User:     "root",
        Password: "password",
        Name:     "database_migration",
    }
    dbDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbConfig.User, dbConfig.Password,
        dbConfig.Host, dbConfig.Port, dbConfig.Name)
    
    sqlxDB, err := sqlx.Connect("mysql", dbDsn)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer sqlxDB.Close()

    // Set up Fiber
    app := fiber.New()
    app.Use(cors.New())

    // Define the migration directory
    migrationDir := "migrations"
    if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
        log.Fatalf("Migration directory does not exist: %s", migrationDir)
    }

    // Register routes
    app.Get("/migrate", func(c *fiber.Ctx) error {
        return migrate(c, sqlxDB, migrationDir)
    })

    // Start the server
    log.Println("Starting Fiber server on :3000")
    if err := app.Listen(":3000"); err != nil {
        log.Fatalf("Error starting Fiber server: %v", err)
    }
}

// migrate performs the migration operation
func migrate(c *fiber.Ctx, db *sqlx.DB, migrationDir string) fiber.fiberError {
    paths := []string{migrationDir}
    conf := goose.DBConf{
        Driver:       "mysql",
        OpenString:  db.DB.String(),
        Diag:        goose.EnhancedReporter{
            Migrations:       &goose.Config{
                TableName: "goose_db_version",
            TZ:         "Local",
            Location:   goose.RunningInContainer()},
        },
        EnvPrefix:   "", // Set if you have an env prefix for goose.
    }
    err := goose.SetDialect(conf.Driver, &conf)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to set dialect",
        })
    }

    err = goose.Up(conf, paths, "head")
    if err != nil {
        log.Printf("Error while applying migrations: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to apply migrations",
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Migration successful",
    })
}

// RunningInContainer returns true if the application is running in a container
func goose.RunningInContainer() bool {
    // Implement logic to determine if application is running in a container
    return false
}
