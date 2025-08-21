// 代码生成时间: 2025-08-21 23:25:18
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/mysql"
    "github.com/golang-migrate/migrate/v4/source/file"
    "golang.org/x/crypto/ssh/terminal"
)

// DatabaseConfig holds the configuration for the database connection
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
}

// MigrateUp migrates the database to the latest version
func MigrateUp(config DatabaseConfig) error {
    var sourceURL string
    if config.Port == 3306 { // Assuming MySQL
        sourceURL = fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local",
            config.User,
            config.Password,
            config.Host,
            config.Port,
            config.Name,
        )
    } else {
        return fmt.Errorf("unsupported database port: %d", config.Port)
    }

    m, err := migrate.NewWithSourceInstance("file://migrations", sourceURL, &mysql.Config{})
    if err != nil {
        return err
    }

    err = m.Up()
    if err != nil {
        if err == migrate.ErrNoChange {
            fmt.Println("No migrations to apply.")
            return nil
        }
        return err
    }
    fmt.Println("Migration successful.")
    return nil
}

// MigrateDown rolls back the last migration
func MigrateDown(config DatabaseConfig) error {
    var sourceURL string
    if config.Port == 3306 {
        sourceURL = fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local",
            config.User,
            config.Password,
            config.Host,
            config.Port,
            config.Name,
        )
    } else {
        return fmt.Errorf("unsupported database port: %d", config.Port)
    }

    m, err := migrate.NewWithSourceInstance("file://migrations", sourceURL, &mysql.Config{})
    if err != nil {
        return err
    }

    err = m.Down()
    if err != nil {
        return err
    }
    fmt.Println("Migration rolled back.")
    return nil
}

func main() {
    config := DatabaseConfig{
        Host:     os.Getenv("DB_HOST"),
        Port:     3306,
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        Name:     os.Getenv("DB_NAME"),
    }

    if config.User == "" || config.Password == "" || config.Name == "" {
        fmt.Println("Please set the database credentials in the environment variables.")
        os.Exit(1)
    }

    // Prompt user for password if not set in environment variables
    if config.Password == "" {
        fmt.Print("Enter database password: ")
        bytePassword, _ := terminal.ReadPassword(0)
        fmt.Println()
        config.Password = string(bytePassword)
    }

    // Migrate up
    if err := MigrateUp(config); err != nil {
        log.Fatal(err)
    }
}
