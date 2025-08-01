// 代码生成时间: 2025-08-01 18:20:24
package main

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "time"

    "github.com/gofiber/fiber/v2"
)

// MemoryUsage contains the memory statistics
type MemoryUsage struct {
    Alloc          uint64 `json:"alloc"`          // bytes allocated and not yet freed
    TOTAL_ALLOC   uint64 `json:"total_alloc"`    // bytes allocated (even if freed)
    Mallocs       uint64 `json:"mallocs"`       // number of mallocs
    Frees         uint64 `json:"frees"`         // number of frees
    HeapAlloc     uint64 `json:"heap_alloc"`     // bytes allocated in heap
    HeapSys       uint64 `json:"heap_sys"`       // heap overhead
    HeapIdle      uint64 `json:"heap_idle"`      // heap that is idle
    HeapInuse     uint64 `json:"heap_inuse"`     // heap that is in use
    HeapReleased  uint64 `json:"heap_released"`  // bytes released to the OS
    HeapObjects   uint64 `json:"heap_objects"`   // number of allocated heap objects
    StackInuse    uint64 `json:"stack_inuse"`    // bytes in use by the stack allocator
    StackSys      uint64 `json:"stack_sys"`      // stack allocator overhead
    MSpanInuse    uint64 `json:"mspan_inuse"`    // mspan structures in use
    MSpanSys      uint64 `json:"mspan_sys"`      // mspan structures overhead
    Sys           uint64 `json:"sys"`           // total bytes of memory obtained by system
    Lookback      uint64 `json:"lookback"`       // GC trigger
    Hooks         [2]uint64 `json:"hooks"`       // profiling/tracing hooks
    NextGC        uint64 `json:"next_gc"`       // next GC byte threshold
    LastGC        uint64 `json:"last_gc"`       // last GC completed timestamp
    Pausing      struct {
        TOTAL    uint64 `json:"total"`    // total pause time
        NUM_GC   uint64 `json:"num_gc"`   // number of GC completed
    } `json:"pausing"`
    NumGC         uint64 `json:"num_gc"`         // number of GC completed
}

// GetMemoryUsage retrieves memory usage statistics
func GetMemoryUsage() MemoryUsage {
    m := &runtime.MemStats{}
    runtime.ReadMemStats(m)

    return MemoryUsage{
        Alloc:          m.Alloc,
        TOTAL_ALLOC:    m.TotalAlloc,
        Mallocs:       m.Mallocs,
        Frees:         m.Frees,
        HeapAlloc:     m.HeapAlloc,
        HeapSys:       m.HeapSys,
        HeapIdle:      m.HeapIdle,
        HeapInuse:     m.HeapInuse,
        HeapReleased:  m.HeapReleased,
        HeapObjects:   m.HeapObjects,
        StackInuse:    m.StackInuse,
        StackSys:      m.StackSys,
        MSpanInuse:    m.MSpanInuse,
        MSpanSys:      m.MSpanSys,
        Sys:           m.Sys,
        Lookback:      m.Lookups,
        Hooks:         m.PauseNs,
        NextGC:        m.NextGC,
        LastGC:        m.LastGC,
        Pausing: struct{
            TOTAL: m.PauseTotalNs,
            NUM_GC: m.NumGC,
        },
        NumGC:         m.NumGC,
    }
}

// setupRoutes sets up the routes for the application
func setupRoutes(app *fiber.App) {
    // GET /memory-usage
    app.Get("/memory-usage", func(c *fiber.Ctx) error {
        usage := GetMemoryUsage()
        return c.JSON(usage)
    })
}

func main() {
    // Initialize Fiber app
    app := fiber.New()

    // Setup routes
    setupRoutes(app)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
