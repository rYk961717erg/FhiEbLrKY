// 代码生成时间: 2025-09-20 21:36:14
package main

import (
    "context"
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sync"

    "github.com/gofiber/fiber/v2"
)

// FileInfo represents the metadata of a file
type FileInfo struct {
    Name     string `json:"name"`
    Size     int64  `json:"size"`
    ModTime  string `json:"modTime"`
    IsDir    bool   `json:"isDir"`
    ErrorMsg string `json:"errorMsg"`
}

// BackupSyncConfig contains the configuration for the backup and sync tool
type BackupSyncConfig struct {
    SourceDir  string `json:"sourceDir"`
    TargetDir  string `json:"targetDir"`
    IgnoreDirs []string `json:"ignoreDirs"`
}

// App is the main application struct
type App struct {
    config BackupSyncConfig
    mu     sync.Mutex
}

// NewApp creates a new instance of the App
func NewApp(config BackupSyncConfig) *App {
    return &App{
        config: config,
    }
}

// BackupAndSync performs the backup and sync operation
func (a *App) BackupAndSync(ctx context.Context) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    sourceInfo, err := os.Stat(a.config.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to stat source dir: %w", err)
    }
    if !sourceInfo.IsDir() {
        return fmt.Errorf("source path is not a directory")
    }

    targetInfo, err := os.Stat(a.config.TargetDir)
    if err != nil {
        return fmt.Errorf("failed to stat target dir: %w", err)
    }
    if !targetInfo.IsDir() {
        return fmt.Errorf("target path is not a directory")
    }

    err = filepath.WalkDir(a.config.SourceDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
