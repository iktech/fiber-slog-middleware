package logger

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time"
)

// New creates a new middleware handler
func New(component string) fiber.Handler {
	// Set variables
	var (
		Start time.Time
		Stop  time.Time
	)

	// Return new handler
	return func(c *fiber.Ctx) error {
		Start = time.Now()

		if component == "" {
			component = "logger"
		}

		// Handle request, store err to return
		chainErr := c.Next()

		Stop = time.Now()
		slog.Info("request", "component", component, "method", c.Method(), "response_time", Stop.Sub(Start).Round(time.Millisecond),
			"source_ip", c.IP(), "status", c.Response().StatusCode(), "path", c.Path())
		// End chain
		return chainErr
	}
}
