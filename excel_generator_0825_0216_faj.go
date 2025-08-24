// 代码生成时间: 2025-08-25 02:16:20
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/xuri/excelize/v2"
    "github.com/gofiber/fiber/v2"
)

// ExcelGenerator 结构体，用于存储Excel文件生成器的状态
type ExcelGenerator struct {
    // 不包含任何字段，仅作为一个包含方法的结构体
}

// NewExcelGenerator 初始化Excel生成器
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{}
}

// GenerateExcel 生成Excel文件
func (e *ExcelGenerator) GenerateExcel(data [][]string, sheetName string) (string, error) {
    file, err := excelize.CreateFile()
    if err != nil {
        return "", err
    }
    defer file.Close()
    // 设置Excel文件名
    fileName := fmt.Sprintf("%s_%s.xlsx", sheetName, time.Now().Format("2006-01-02_150405"))
    // 设置工作表名称
    if err := file.NewSheet(sheetName); err != nil {
        return "", err
    }
    // 设置文件的默认工作表
    if err := file.SetActiveSheet(file.GetSheetIndex(sheetName)); err != nil {
        return "