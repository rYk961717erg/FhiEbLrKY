// 代码生成时间: 2025-08-08 01:11:42
package main

import (
    "fmt"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cache"
)

// CacheService 结构体包含缓存相关的配置和方法
type CacheService struct {
    cache *cache.Cache
}

// NewCacheService 创建一个新的 CacheService 实例
func NewCacheService() *CacheService {
    return &CacheService{
        cache: cache.New(
            cache.Settings{
                Expiration: 5 * time.Minute, // 设置缓存过期时间为5分钟
                CalculateExpiration: func(c *cache.Cache, ctx *fiber.Ctx) time.Duration {
                    return c.Settings.Expiration // 使用默认过期时间
                },
            },
        ),
    }
}

// SetCache 设置缓存项
func (c *CacheService) SetCache(ctx *fiber.Ctx, key string, value interface{}) error {
    err := c.cache.Set(key, value, -1) // -1 表示使用默认过期时间
    if err != nil {
        return fmt.Errorf("failed to set cache: %w", err)
    }
    return nil
}

// GetCache 获取缓存项
func (c *CacheService) GetCache(ctx *fiber.Ctx, key string) (interface{}, error) {
    cachedValue, err := c.cache.Get(key)
    if err != nil {
        return nil, fmt.Errorf("failed to get cache: %w", err)
    }
    return cachedValue, nil
}

// main 函数启动 Fiber 服务并配置缓存中间件
func main() {
    app := fiber.New()

    // 实例化缓存服务
    cacheService := NewCacheService()

    // 使用 Fiber 的缓存中间件
    app.Use(cache.New(cache.Settings{Expiration: 5 * time.Minute}))

    // 缓存测试路由
    app.Get("/cache", func(ctx *fiber.Ctx) error {
        key := "testKey"
        value, err := cacheService.GetCache(ctx, key)
        if err != nil {
            if err.Error() == "not found" { // 缓存未命中
                value = "new value"
                // 将新值设置到缓存
                if setErr := cacheService.SetCache(ctx, key, value); setErr != nil {
                    return setErr
                }
            } else {
                return err
            }
        }
        return ctx.SendString(fmt.Sprintf("Cache value: %v", value))
    })

    // 启动服务
    if err := app.Listen(":3000"); err != nil {
        fmt.Printf("Error starting server: %s
", err)
    }
}