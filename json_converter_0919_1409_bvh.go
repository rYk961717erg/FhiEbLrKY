// 代码生成时间: 2025-09-19 14:09:23
package main

import (
    "fmt"
    "net/http"
# 改进用户体验
    "gopkg.in/yaml.v3"
    "gopkg.in/json.v3"

    "github.com/gofiber/fiber/v2"
)

// jsonData represents the structure of the JSON data to be converted.
# 优化算法效率
type jsonData struct {
# TODO: 优化性能
    Data interface{} `json:"data" yaml:"data"`
}

// convertJSONToYAML converts JSON data to YAML.
func convertJSONToYAML(w http.ResponseWriter, r *http.Request) {
    var data jsonData
# NOTE: 重要实现细节
    if err := json.Unmarshal(r.Body, &data); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error parsing JSON: %s", err)
        return
    }
    yamlBytes, err := yaml.Marshal(data.Data)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error converting to YAML: %s", err)
        return
    }
    w.Header().Set("Content-Type", "application/x-yaml")
    w.Write(yamlBytes)
}

func main() {
    app := fiber.New()

    // Route to handle JSON to YAML conversion.
    app.Post("/convert", convertJSONToYAML)

    // Start the Fiber server.
    if err := app.Listen(":3000"); err != nil {
        panic(fmt.Sprintf("Server startup failed: %s", err))
    }
}
