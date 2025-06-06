package security

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// Config holds the security configuration
type Config struct {
	BodyLimit       int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	RateLimitMax    int
	RateLimitWindow time.Duration
}

// DefaultConfig returns the default security configuration
func DefaultConfig() *Config {
	return &Config{
		BodyLimit:       1 * 1024 * 1024, // 1MB
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     120 * time.Second,
		RateLimitMax:    10,
		RateLimitWindow: 1 * time.Minute,
	}
}

// ApplySecurityMiddleware applies all security middleware to the app
func ApplySecurityMiddleware(app *fiber.App, config *Config) {
	// Global middleware
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(compress.New())

	// Rate limiting configuration
	limiterConfig := limiter.Config{
		Max:        config.RateLimitMax,
		Expiration: config.RateLimitWindow,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() + ":" + c.Get("User-Agent")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "Rate limit exceeded. Please try again later.",
				"retry_after": "1 minute",
			})
		},
	}

	app.Use(limiter.New(limiterConfig))

	// Security headers
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		return c.Next()
	})
}
