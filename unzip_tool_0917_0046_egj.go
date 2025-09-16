// 代码生成时间: 2025-09-17 00:46:12
package main

import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// UnzipHandler is a Fiber handler function that handles file uploads and unzips them.
func UnzipHandler(c *fiber.Ctx) error {
    // Check if the file is in the request
    if !c.HasFile("") {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "No file uploaded",
        })
    }

    // Get the file from the request
    file, err := c.FormFile("")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    defer file.Close()

    // Define the destination folder for the unzipped files
    dest := "unzipped/" + strings.ReplaceAll(file.Filename, ".", "_")
    if err := os.MkdirAll(dest, 0755); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Open the zip file
    rc, err := zip.OpenReader(file.Path())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    defer rc.Close()

    // Loop through the files in the zip
    for _, f := range rc.File {
        // Create the full path to the file
        filePath := filepath.Join(dest, f.Name)

        // Check for zip entry that is a directory and create it
        if f.FileInfo().IsDir() {
            if err := os.MkdirAll(filePath, f.Mode()); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "error": err.Error(),
                })
            }
            continue
        }

        // Create the file
        outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        defer outFile.Close()

        // Copy the file content from the zip to the new file
        rcFile, err := f.Open()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        defer rcFile.Close()

        _, err = io.Copy(outFile, rcFile)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
    }

    // Return a success message
    return c.JSON(fiber.Map{
        "message": "File successfully unzipped",
        "destination": dest,
    })
}

func main() {
    // Create a new Fiber app
    app := fiber.New()

    // Set the file upload limit to 10MB
    app.SetAppSettings(fiber.AppSettings{
        UploadFileLimit: 10 * 1024 * 1024,
    })

    // Define the route for file uploads
    app.Post("/unzip", UnzipHandler)

    // Start the server
    log.Fatal(app.Listen(":8080"))
}
