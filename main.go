package main

// this will be used as router as well

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/manohySr/weather-api/cache"
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

	app := fiber.New()

	app.Get("/weather/:city", func(c *fiber.Ctx) error {
		city := strings.ToLower(c.Params("city"))

		log.Println("GET /weather/", city)
		res, err := weatherService.GetWeather(city)
		if err != nil {
			log.Println("Failed to get weather:", err)

			// Check if it's a "city not found" error
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
