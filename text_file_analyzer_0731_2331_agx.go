// 代码生成时间: 2025-07-31 23:31:09
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// TextAnalyzer contains methods for analyzing text files.
type TextAnalyzer struct{}

// AnalyzeTextFile analyzes the content of a text file and returns statistics.
func (a *TextAnalyzer) AnalyzeTextFile(filePath string) (map[string]int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines, words, bytes int
    for scanner.Scan() {
        line := scanner.Text()
        lines++
        words += len(strings.Fields(line))
        bytes += len(line)
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return map[string]int{
        "lines": lines,
        "words": words,
        "bytes": bytes,
    }, nil
}

func main() {
    app := fiber.New()

    // Define the route for analyzing a text file.
    app.Get("/analyze", func(c *fiber.Ctx) error {
        filePath := c.Query("file")
        if filePath == "" {
            return fiber.StatusBadRequest
        }

        analyzer := TextAnalyzer{}
        stats, err := analyzer.AnalyzeTextFile(filePath)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // Return the statistics as JSON.
        return c.JSON(fiber.Map{
            "lines": stats["lines"],
            "words": stats["words"],
            "bytes": stats["bytes"],
        })
    })

    // Start the Fiber server.
    log.Fatal(app.Listen(":3000"))
}
