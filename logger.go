package logger

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time"
)

// New creates a new middleware handler
func New(component string) fiber.Handler {
	return NewWithExclusions(component, nil)
}

// NewWithExclusions creates a new middleware handler with list of routes excluded from logging
func NewWithExclusions(component string, exclusions []string) fiber.Handler {
	// Set variables
	var (
		Start  time.Time
		Stop   time.Time
		routes map[string]interface{}
	)

	routes = make(map[string]interface{})
	for _, exclusion := range exclusions {
		routes[exclusion] = true
	}

	// Return new handler
	return func(c *fiber.Ctx) error {
		Start = time.Now()

		if component == "" {
			component = "logger"
		}

		// Handle request, store err to return
		chainErr := c.Next()

		log := true
		if c.Route() != nil {
			if routes[c.Route().Path] != nil {
				log = false
			}
		}

		if log {
			Stop = time.Now()
			slog.Info("request", "component", component, "method", c.Method(), "response_time", Stop.Sub(Start).Round(time.Millisecond),
				"source_ip", c.IP(), "status", c.Response().StatusCode(), "path", c.Path())
		}

		// End chain
		return chainErr
	}
}
