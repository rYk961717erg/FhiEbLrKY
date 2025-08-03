// 代码生成时间: 2025-08-03 20:40:16
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// Theme represents the available themes
type Theme struct {
    Name string `json:"name"`
}

// ThemeService is a struct that manages theme switching
type ThemeService struct {
    // Stores the current theme
    currentTheme *Theme
}

// NewThemeService initializes a new ThemeService instance
func NewThemeService() *ThemeService {
    return &ThemeService{
        currentTheme: &Theme{Name: "light"}, // Default theme
    }
}

// SwitchTheme changes the current theme to the specified theme
func (s *ThemeService) SwitchTheme(themeName string) (*Theme, error) {
    // Check if the theme is valid (for now, we only have 'light' and 'dark')
    if themeName == "light" || themeName == "dark" {
        s.currentTheme = &Theme{Name: themeName}
        return s.currentTheme, nil
    } else {
        // Return an error if the theme is not valid
        return nil, fmt.Errorf("theme '%s' is not supported", themeName)
    }
}

// ThemeHandler is the Fiber handler for theme switching
func ThemeHandler(themeService *ThemeService) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var newTheme Theme
        // Parse the new theme from the request body
        if err := c.BodyParser(&newTheme); err != nil {
            return err
        }
        // Switch theme using the ThemeService
        newTheme, err := themeService.SwitchTheme(newTheme.Name)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        // Return the new theme in the response
        return c.JSON(newTheme)
    }
}

func main() {
    // Initialize Fiber
    app := fiber.New()

    // Initialize ThemeService
    themeService := NewThemeService()

    // Set up the theme switcher route
    app.Post("/api/theme", ThemeHandler(themeService))

    // Start the server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
