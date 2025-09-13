// 代码生成时间: 2025-09-13 11:17:36
package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/gofiber/fiber/v2"
    "log"
)

// DatabaseConfig 配置数据库连接参数
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// DatabaseManager 管理数据库连接池
type DatabaseManager struct {
    db *sql.DB
}

// NewDatabaseManager 创建一个新的数据库管理器
func NewDatabaseManager(cfg DatabaseConfig) (*DatabaseManager, error) {
    // 构建DSN（数据源名称）
    dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
    // 打开数据库连接
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    // 设置连接池参数
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(10)
    db.SetConnMaxLifetime(5 * time.Minute)
    // 测试连接
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DatabaseManager{db: db}, nil
}

// Close 关闭数据库连接池
func (dm *DatabaseManager) Close() error {
    if dm.db != nil {
        return dm.db.Close()
    }
    return nil
}

func main() {
    cfg := DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "root",
        Password: "password",
        DBName:   "database_name",
    }
    // 初始化数据库连接池
    dbManager, err := NewDatabaseManager(cfg)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %s", err)
    }
    defer dbManager.Close()

    // 设置Fiber路由
    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // 启动服务器
    log.Fatal(app.Listen(":3000"))
}
