package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CurrentWeatherResponse struct {
	Datetime   string  `json:"datetime"`   // "14:17:00"
	Temp       float64 `json:"temp"`       // 66
	FeelsLike  float64 `json:"feelslike"`  // 66
	Humidity   float64 `json:"humidity"`   // 60
	Precip     float64 `json:"precip"`     // 0
	PrecipProb float64 `json:"precipprob"` // 0
	WindGust   float64 `json:"windgust"`   // 12.2
	WindSpeed  float64 `json:"windspeed"`  // 6.1
	WindDir    float64 `json:"winddir"`    // 259 (°)
	Pressure   float64 `json:"pressure"`   // 1010 (hPa)
	Visibility float64 `json:"visibility"` // 6.2 (mi / km)
	CloudCover float64 `json:"cloudcover"` // 88 (%)
	UVIndex    float64 `json:"uvindex"`    // 1
	Conditions string  `json:"conditions"` // "Partially cloudy"      // "partly-cloudy-day"
}

type ApiResponse struct {
	CurrentConditions CurrentWeatherResponse `json:"currentConditions"`
}

func GetWeatherCurrentData(city string) (*CurrentWeatherResponse, error) {
	var key string = os.Getenv("key")
	var url string = fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=us&key=%s&contentType=json", city, key)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var weather ApiResponse
	if err := json.NewDecoder(res.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather.CurrentConditions, nil
}
