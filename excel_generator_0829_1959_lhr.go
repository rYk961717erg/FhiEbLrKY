// 代码生成时间: 2025-08-29 19:59:44
package main

import (
    "excelize"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "os"
)

// ExcelGenerator 定义一个Excel生成器结构体
type ExcelGenerator struct {
    // 可以添加更多字段来扩展Excel生成器的功能
}

// NewExcelGenerator 创建一个新的ExcelGenerator实例
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{}
}

// GenerateExcel 生成Excel文件
func (g *ExcelGenerator) GenerateExcel(data [][]string) (string, error) {
    // 创建一个新的Excel文件
    f := excelize.NewFile()

    // 设置Excel文件的标题
    for i, row := range data {
        if i == 0 {
            // 设置标题行
            for j, title := range row {
                f.SetCellValue("Sheet1", fmt.Sprintf("A%d", j+1), title)
            }
        } else {
            // 设置数据行
            for j, value := range row {
                f.SetCellValue("Sheet1", fmt.Sprintf("A%d", j+1+i), value)
            }
        }
    }

    // 保存Excel文件
    filePath := "output.xlsx"
    if err := f.SaveAs(filePath); err != nil {
        return "", err
    }

    return filePath, nil
}

func main() {
    app := fiber.New()

    // 创建Excel生成器实例
    generator := NewExcelGenerator()

    // 定义路由，用于生成Excel文件
    app.Get("/generate", func(c *fiber.Ctx) error {
        // 模拟一些Excel数据
        data := [][]string{
            {"Name", "Age", "City"},
            {"John", "30", "New York"},
            {"Jane", "25", "Los Angeles"},
        }

        // 生成Excel文件
        filePath, err := generator.GenerateExcel(data)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to generate Excel file",
            })
        }

        // 下载Excel文件
        return c.SendFile(filePath)
    })

    // 启动Fiber服务器
    if err := app.Listen(":3000"); err != nil {
        panic(err)
    }
}
