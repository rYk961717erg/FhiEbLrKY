// 代码生成时间: 2025-09-05 15:41:12
package main

import (
    "fmt"
    "net/http"
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/utils"
    "github.com/stretchr/testify/assert"
)

// 测试的路由
const testRoute = "/test"

// TestMain 是测试的入口点，它将启动Fiber应用并运行测试
func TestMain(m *testing.M) {
    app := fiber.New()
    
    // 在这里注册你的路由和中间件
    app.Get(testRoute, func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    
    // 启动Fiber应用
    utils.Ready(app)
    
    // 运行测试
    m.Run()
    
    // 停止Fiber应用
    app.Shutdown()
}

// TestIntegration 是一个集成测试的例子，它测试/test路由
func TestIntegration(t *testing.T) {
    assert := assert.New(t)
    
    // 发送GET请求到/test路由
    resp, err := fibertest.Test(app, http.MethodGet, testRoute)
    assert.NoError(err)
    assert.Equal(http.StatusOK, resp.StatusCode)
    
    // 验证响应体的内容
    body, _ := utils.Read.resp.Body)
    assert.Equal("Hello, World!", string(body))
}
