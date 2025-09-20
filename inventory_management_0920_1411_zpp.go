// 代码生成时间: 2025-09-20 14:11:33
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// InventoryItem represents an item in the inventory
type InventoryItem struct {
    ID          uint   `json:"id"`
    Name        string `json:"name"`
    Quantity    int    `json:"quantity"`
    Description string `json:"description"`
}

// InventoryService handles inventory operations
type InventoryService struct {
    items []InventoryItem
}

// NewInventoryService creates a new inventory service
func NewInventoryService() *InventoryService {
    return &InventoryService{
        items: make([]InventoryItem, 0),
    }
}

// AddItem adds a new item to the inventory
func (s *InventoryService) AddItem(item InventoryItem) (uint, error) {
    s.items = append(s.items, item)
    return item.ID, nil
}

// UpdateItem updates an existing item in the inventory
func (s *InventoryService) UpdateItem(id uint, item InventoryItem) error {
    for i, existingItem := range s.items {
        if existingItem.ID == id {
            s.items[i] = item
            return nil
        }
    }
    return fmt.Errorf("item with id %d not found", id)
}

// DeleteItem removes an item from the inventory by ID
func (s *InventoryService) DeleteItem(id uint) error {
    for i, item := range s.items {
        if item.ID == id {
            s.items = append(s.items[:i], s.items[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("item with id %d not found", id)
}

// GetItem retrieves an item from the inventory by ID
func (s *InventoryService) GetItem(id uint) (InventoryItem, error) {
    for _, item := range s.items {
        if item.ID == id {
            return item, nil
        }
    }
    return InventoryItem{}, fmt.Errorf("item with id %d not found", id)
}

// GetAllItems returns a list of all items in the inventory
func (s *InventoryService) GetAllItems() []InventoryItem {
    return s.items
}

func main() {
    app := fiber.New()
    service := NewInventoryService()

    // Initialize the inventory with some items
    service.AddItem(InventoryItem{ID: 1, Name: "Apple", Quantity: 100, Description: "Fresh apples"})
    service.AddItem(InventoryItem{ID: 2, Name: "Banana", Quantity: 150, Description: "Ripe bananas"})

    // Routes
    app.Get("/items", func(c *fiber.Ctx) error {
        items := service.GetAllItems()
        return c.JSON(items)
    })

    app.Get("/items/:id", func(c *fiber.Ctx) error {
        id := uint(c.Params("id"))
        item, err := service.GetItem(id)
        if err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(item)
    })

    app.Post("/items", func(c *fiber.Ctx) error {
        var item InventoryItem
        if err := c.BodyParser(&item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        _, err := service.AddItem(item)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(item)
    })

    app.Put("/items/:id", func(c *fiber.Ctx) error {
        id := uint(c.Params("id"))
        var item InventoryItem
        if err := c.BodyParser(&item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        if err := service.UpdateItem(id, item); err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(item)
    })

    app.Delete("/items/:id", func(c *fiber.Ctx) error {
        id := uint(c.Params("id"))
        if err := service.DeleteItem(id); err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendStatus(fiber.StatusNoContent)
    })

    log.Fatal(app.Listen(":3000"))
}
