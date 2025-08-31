// 代码生成时间: 2025-08-31 08:13:59
package main

import (
    "fmt"
    "net/http"
    "os"
    "runtime"
    "time"

    "github.com/gofiber/fiber/v2"
)

// MemoryUsage provides memory usage statistics
type MemoryUsage struct {
    Alloc       uint64 `json:"alloc"`       // bytes allocated and not yet freed
    TotAlloc    uint64 `json:"tot_alloc"`    // bytes allocated (even if freed)
    Sys         uint64 `json:"sys"`         // bytes obtained from system (includes stack)
   Mallocs     uint64 `json:"mallocs"`     // times memory was allocated
   Frees       uint64 `json:"frees"`       // times memory was freed
    HeapAlloc   uint64 `json:"heap_alloc"`   // bytes allocated in heap
    HeapSys     uint64 `json:"heap_sys"`     // heap system bytes
    HeapIdle    uint64 `json:"heap_idle"`    // heap idle bytes
    HeapInuse   uint64 `json:"heap_inuse"`   // heap in-use bytes
    HeapReleased uint64 `json:"heap_released"` // bytes released to the OS
    HeapObjects uint64 `json:"heap_objects"` // total number of allocated heap objects
}

// getMemoryUsage returns the current memory usage statistics
func getMemoryUsage() MemoryUsage {
    m := &runtime.MemStats{}
    runtime.ReadMemStats(m)
    return MemoryUsage{
        Alloc:       m.Alloc,
        TotAlloc:    m.TotalAlloc,
        Sys:         m.Sys,
        Mallocs:     m.Mallocs,
        Frees:       m.Frees,
        HeapAlloc:   m.HeapAlloc,
        HeapSys:     m.HeapSys,
        HeapIdle:    m.HeapIdle,
        HeapInuse:   m.HeapInuse,
        HeapReleased: m.HeapReleased,
        HeapObjects: m.HeapObjects,
    }
}

// MemoryUsageHandler returns the current memory usage statistics
func MemoryUsageHandler(c *fiber.Ctx) error {
    memUsage := getMemoryUsage()
    return c.JSON(memUsage)
}

func main() {
    app := fiber.New()

    // Register the memory usage handler
    app.Get("/memory", MemoryUsageHandler)

    // Start the server
    addr := ":3000"
    fmt.Printf("Server is running on %s", addr)
    if err := app.Listen(addr); err != nil {
        fmt.Println("Error starting server: ", err)
        os.Exit(1)
    }
}
