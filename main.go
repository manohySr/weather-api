package main

// this will be used as router as well

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/manohySr/weather-api/weather"
)

func main() {
	log.Println("Starting...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	res, err := weather.GetWeatherCurrentData("london")
	if err != nil {
		fmt.Println("Error fetching")
		return
	}
	jsonBytes, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Print("Error marshalling to JSON: ", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
