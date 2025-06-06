package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/manohySr/weather-api/cache"
	"github.com/manohySr/weather-api/security"
	"github.com/manohySr/weather-api/weather"
)

func main() {
	log.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisClient := cache.InitRedis()
	if redisClient == nil {
		log.Println("warning: Running without redis cache")
	}

	weatherService := weather.NewWeatherService(redisClient)

	// Create app with security config
	config := fiber.Config{
		BodyLimit:             security.DefaultConfig().BodyLimit,
		ReadTimeout:           security.DefaultConfig().ReadTimeout,
		WriteTimeout:          security.DefaultConfig().WriteTimeout,
		IdleTimeout:           security.DefaultConfig().IdleTimeout,
		DisableStartupMessage: true,
	}
	app := fiber.New(config)

	// Apply security middleware with default config
	security.ApplySecurityMiddleware(app, security.DefaultConfig())

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		duration := time.Since(start)
		log.Printf("📝 Request => %s %s | Status: %d | Duration: %v", c.Method(), c.Path(), c.Response().StatusCode(), duration)
		return c.Next()
	})

	app.Get("/weather/:city", func(c *fiber.Ctx) error {
		city := strings.ToLower(c.Params("city"))

		res, err := weatherService.GetWeather(city)
		if err != nil {
			log.Printf("[%s] Failed to get weather: %v", c.Get("X-Request-ID"), err)

			if strings.Contains(err.Error(), "not found") {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "City not found",
					"city":  city,
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable to fetch weather data",
				"city":  city,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
			"city": city,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server is running at http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
