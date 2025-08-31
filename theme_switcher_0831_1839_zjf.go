// 代码生成时间: 2025-08-31 18:39:27
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// Theme represents the theme structure
type Theme struct {
    Name string `json:"name"`
}

// ThemeService handles theme switching
type ThemeService struct {
    themes []Theme
}

// NewThemeService creates a new ThemeService with predefined themes
func NewThemeService() *ThemeService {
    return &ThemeService{
        themes: []Theme{
            {Name: "light"},
            {Name: "dark"},
        },
    }
}

// SwitchTheme changes the theme based on the provided theme name
func (s *ThemeService) SwitchTheme(name string) (Theme, error) {
    for _, theme := range s.themes {
        if theme.Name == name {
            return theme, nil
        }
    }
    return Theme{}, fmt.Errorf("theme not found: %s", name)
}

func main() {
    app := fiber.New()
    themeService := NewThemeService()

    // Endpoint to switch theme
    app.Post("/switch-theme", func(c *fiber.Ctx) error {
        var themeName string
        if err := c.BodyParser(&themeName); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "invalid request body",
            })
        }

        newTheme, err := themeService.SwitchTheme(themeName)
        if err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "message": "theme switched successfully",
            "newTheme": newTheme,
        })
    })

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
