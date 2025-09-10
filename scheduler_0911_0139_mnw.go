// 代码生成时间: 2025-09-11 01:39:35
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/robfig/cron/v3"
)

// Scheduler struct to hold the cron scheduler
type Scheduler struct {
    Cron *cron.Cron
}

// NewScheduler creates a new scheduler with given schedule
func NewScheduler(schedule string, task func()) *Scheduler {
    cron := cron.New()
    cron.AddFunc(schedule, task)
    cron.Start()
    return &Scheduler{Cron: cron}
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.Cron.Stop()
}

// StartFiber starts a Fiber web server with a route to trigger the task
func StartFiber(scheduler *Scheduler, task func()) *fiber.App {
    app := fiber.New()
    app.Get("/trigger", func(c *fiber.Ctx) error {
        task()
        return c.SendStatus(fiber.StatusOK)
    })
    return app
}

func main() {
    // Define a task that will be executed by the scheduler
    task := func() {
        fmt.Println("Task executed at: ", time.Now().Format(time.RFC1123))
    }

    // Create a new scheduler with a schedule (e.g., every 5 seconds)
    scheduler := NewScheduler("*/5 * * * *", task)
    defer scheduler.Stop()

    // Start Fiber web server
    app := StartFiber(scheduler, task)
    app.Listen(":3000")
}
