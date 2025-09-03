// 代码生成时间: 2025-09-03 15:53:55
package main

import (
    "flag"
    "fmt"
    "image"
    "image/jpeg"
# 优化算法效率
    "io/ioutil"
# 优化算法效率
    "os"
    "path/filepath"
# 增强安全性
    "log"
    "github.com/golang/freetype/truetype"
    "github.com/golang/freetype"
    "github.com/nfnt/resize"
    "github.com/gofiber/fiber/v2"
)

// ImageResizer 结构体用于保存图片尺寸调整的配置
type ImageResizer struct {
    Width, Height int
}

// NewImageResizer 创建并返回一个ImageResizer实例
func NewImageResizer(width, height int) *ImageResizer {
    return &ImageResizer{Width: width, Height: height}
}
# TODO: 优化性能

// ResizeImage 调整图片尺寸
func (r *ImageResizer) ResizeImage(img image.Image) image.Image {
    return resize.Resize(uint(r.Width), uint(r.Height), img, resize.Lanczos3)
}

// ProcessImageFolder 处理文件夹内的所有图片
func (r *ImageResizer) ProcessImageFolder(folderPath string) error {
    files, err := ioutil.ReadDir(folderPath)
    if err != nil {
        return err
