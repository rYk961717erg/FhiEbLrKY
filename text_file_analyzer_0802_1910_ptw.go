// 代码生成时间: 2025-08-02 19:10:29
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// TextAnalyzer struct to hold file content
type TextAnalyzer struct {
    Content string
}

// AnalyzeContent analyzes the text file content
func (t *TextAnalyzer) AnalyzeContent() (map[string]int, error) {
    // Create a map to hold word counts
    wordCounts := make(map[string]int)

    // Split the content into words
    words := strings.Fields(t.Content)

    for _, word := range words {
        // Increment the count for each word
        wordCounts[strings.ToLower(word)]++
    }

    return wordCounts, nil
}

// AnalyzeFileContent handles the HTTP request to analyze file content
func AnalyzeFileContent(c *fiber.Ctx) error {
    // Read the file content from the request
    file, err := c.FormFile("file")
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(err)
    }

    // Read the file content into a string
    content, err := ioutil.ReadFile(file.Path)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(err)
    }

    // Create a TextAnalyzer instance
    analyzer := TextAnalyzer{Content: string(content)}

    // Analyze the file content
    wordCounts, err := analyzer.AnalyzeContent()
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(err)
    }

    // Return the word counts as JSON response
    return c.JSON(wordCounts)
}

func main() {
    // Create a new Fiber instance
    app := fiber.New()

    // Define the route for analyzing file content
    app.Post("/analyze", AnalyzeFileContent)

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}
