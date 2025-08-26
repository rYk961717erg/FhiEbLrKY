// 代码生成时间: 2025-08-26 19:42:28
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/sirupsen/logrus"
)

// LogParser defines the structure for parsing log files
type LogParser struct {
    // Directory is the path to the directory containing log files
    Directory string
    // Pattern specifies the pattern for log file names
    Pattern string
    // Fields defines the fields to parse from each log line
    Fields []string
}

// NewLogParser creates a new instance of LogParser with default settings
func NewLogParser(directory, pattern string, fields []string) *LogParser {
    return &LogParser{
        Directory: directory,
        Pattern:   pattern,
        Fields:    fields,
    }
}

// ParseLogFile parses a single log file and extracts relevant fields
func (p *LogParser) ParseLogFile(filename string) ([]map[string]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var records []map[string]string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        record := make(map[string]string)
        fields := strings.Fields(line)
        for i, field := range fields {
            if i < len(p.Fields) {
                record[p.Fields[i]] = field
            }
        }
        records = append(records, record)
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return records, nil
}

// ParseDirectory parses all log files in the directory matching the pattern
func (p *LogParser) ParseDirectory() ([]map[string][]map[string]string, error) {
    files, err := filepath.Glob(filepath.Join(p.Directory, p.Pattern))
    if err != nil {
        return nil, err
    }

    var results []map[string][]map[string]string
    for _, file := range files {
        records, err := p.ParseLogFile(file)
        if err != nil {
            logrus.WithError(err).WithField("file", file).Error("Failed to parse log file")
            continue
        }
        results = append(results, map[string][]map[string]string{
            filepath.Base(file): records,
        })
    }
    return results, nil
}

func main() {
    app := fiber.New()

    // Log directory and file pattern
    logDirectory := "./logs"
    logPattern := "*.log"
    logFields := []string{"timestamp", "level", "message"}

    parser := NewLogParser(logDirectory, logPattern, logFields)

    // API endpoint to parse log files
    app.Get("/log", func(c *fiber.Ctx) error {
        results, err := parser.ParseDirectory()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to parse log files",
            })
        }
        return c.JSON(results)
    })

    logrus.Info("Starting log parser API on :3000")
    if err := app.Listen(":3000"); err != nil {
        logrus.WithError(err).Fatal("Failed to start server")
    }
}