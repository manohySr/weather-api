package main

// this will be used as router as well

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/manohySr/weather-api/cache"
	"github.com/manohySr/weather-api/weather"
)

func main() {
	log.Println("Starting point")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	redisClient := cache.InitRedis()
	if redisClient == nil {
		log.Println("Warning: Running without Redis cache")
	}
	weatherService := weather.NewWeatherService(redisClient)
	res, err := weatherService.GetWeather("london")

	if err != nil {
		fmt.Println("error while fetching in main.go")
	}

	jsonBytes, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Print("Error marshalling to JSON: ", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
