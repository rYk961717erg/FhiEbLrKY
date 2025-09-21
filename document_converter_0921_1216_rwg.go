// 代码生成时间: 2025-09-21 12:16:14
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/unidoc/unioffice/document"
)

// DocumentConverterHandler handles the document conversion
func DocumentConverterHandler(c *fiber.Ctx) error {
    file, err := c.FormFile("file")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "No file uploaded",
        })
    }
    defer file.Close()
    
    // Save the uploaded file to a temporary location
    tempFile, err := os.CreateTemp(os.TempDir(), "*.docx")
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Unable to create temporary file",
        })
    }
    defer tempFile.Close()
    
    // Copy the file content to the temporary file
    _, err = tempFile.Write(file.Data())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Unable to write to temporary file",
        })
    }

    // Convert the document to PDF
    err = document.Convert(tempFile.Name(), "pdf", c.Query("output"))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Unable to convert document",
        })
    }

    // Return the converted file path in the response
    return c.JSON(fiber.Map{
        "message": "Document converted successfully",
        "outputFile": filepath.Base(c.Query("output")),
    })
}

// main is the entry point of the application
func main() {
    app := fiber.New()

    // Set up document conversion route
    app.Post("/convert", DocumentConverterHandler)

    // Start the Fiber server
    app.Listen(":3000")
}
