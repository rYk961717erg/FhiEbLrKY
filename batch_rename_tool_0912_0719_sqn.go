// 代码生成时间: 2025-09-12 07:19:11
package main

import (
    "flag"
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// RenameOperation defines the operation to be performed on a file
type RenameOperation struct {
    // OldName is the original file name
    OldName string
    // NewName is the new file name
    NewName string
}

// RenameFile renames a file by creating a RenameOperation and performing the actual rename
func RenameFile(operation RenameOperation) error {
    src := operation.OldName
    dest := operation.NewName
    
    // Check if the source file exists
    if _, err := os.Stat(src); os.IsNotExist(err) {
        return fmt.Errorf("source file does not exist: %s", src)
    }
    
    // Perform the rename operation
    if err := os.Rename(src, dest); err != nil {
        return fmt.Errorf("failed to rename file %s to %s: %s", src, dest, err)
    }
    
    return nil
}

// renameFiles takes a slice of RenameOperation and renames each file accordingly
func renameFiles(operations []RenameOperation) error {
    for _, op := range operations {
        if err := RenameFile(op); err != nil {
            return err
        }
    }
    return nil
}

// setupRoutes sets up the routes for the batch rename tool
func setupRoutes(app *fiber.App) {
    app.Post("/rename", func(c *fiber.Ctx) error {
        // Decode the JSON payload into a slice of RenameOperation
        var operations []RenameOperation
        if err := c.BodyParser(&operations); err != nil {
            return fiber.NewError(fiber.StatusUnprocessableEntity, "failed to parse request body")
        }
        
        // Perform the rename operations
        if err := renameFiles(operations); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "failed to rename files")
        }
        
        return c.JSON(fiber.Map{
            "status": "success",
            "message": "files renamed successfully",
        })
    })
}

func main() {
    app := fiber.New()
    setupRoutes(app)
    
    // Start the Fiber web server
    log.Fatal(app.Listen(":3000"))
}

// Note: The above code assumes that the file operations are performed in a directory
// that the application has permission to write to. The actual implementation may require
// additional error handling and validation depending on the use case.
