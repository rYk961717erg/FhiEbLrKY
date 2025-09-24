// 代码生成时间: 2025-09-24 08:02:56
package main

import (
    "fmt"
    "os"
    "log"
    "github.com/360EntSecGroup-Skylar/excelize/v2"
    "github.com/gofiber/fiber/v2"
)

// ExcelGeneratorHandler 处理生成Excel表格的请求
func ExcelGeneratorHandler(c *fiber.Ctx) error {
    // 创建一个新的Excel工作簿
    f := excelize.NewFile()
   defer f.Close()
    
    // 创建一个名为"Sheet1"的工作表
    index := f.NewSheet("Sheet1")
    
    // 设置单元格的值
    f.SetCellValue("Sheet1", "A1", "Name")
    f.SetCellValue("Sheet1", "B1", "Age")
    
    // 写入一些示例数据
    f.SetCellValue("Sheet1", "A2", "John Doe")
    f.SetCellValue("Sheet1", "B2", 30)
    
    // 将Excel文件保存到服务器上
    filePath := "./example.xlsx"
    if err := f.SaveAs(filePath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to save Excel file.",
            "message": err.Error(),
        })
    }
    
    // 将Excel文件发送给客户端
    return c.SendFile(filePath)
}

func main() {
    // 创建Fiber实例
    app := fiber.New()
    
    // 设置API路由
    app.Get("/generate-excel", ExcelGeneratorHandler)
    
    // 启动服务器
    if err := app.Listen(":3000"); err != nil {
        log.Fatal(err)
    }
}
