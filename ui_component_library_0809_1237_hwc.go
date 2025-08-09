// 代码生成时间: 2025-08-09 12:37:50
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// UIComponent represents a generic user interface component
type UIComponent struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

// NewUIComponent creates a new instance of UIComponent
func NewUIComponent(name, version string) *UIComponent {
    return &UIComponent{Name: name, Version: version}
}

// RegisterComponents registers UI components with their details
func RegisterComponents(app *fiber.App) {
    // Define the components
    btn := NewUIComponent("Button", "1.0.0")
    slider := NewUIComponent("Slider\, "1.0.1")
    
    // Handle GET request for components
    app.Get("/components", func(c *fiber.Ctx) error {
        components := []*UIComponent{btn, slider}
        return c.JSON(components)
    })
    
    // Handle GET request for a specific component
    app.Get("/components/:name", func(c *fiber.Ctx) error {
        name := c.Params("name")
        for _, component := range components {
            if component.Name == name {
                return c.JSON(component)
            }
        }
        // Return an error if the component is not found
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": fmt.Sprintf("Component '%s' not found", name),
        })
    })
}

func main() {
    app := fiber.New()
    RegisterComponents(app)
    
    // Start the server
    if err := app.Listen(":3000"); err != nil {
        fmt.Printf("Error starting server: %v", err)
    }
}
