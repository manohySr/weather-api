package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	cacheKeyPrefix = "weather_cache:"
	cacheDuration  = 30 * time.Minute
)

type WeatherService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewWeatherService(redisClient *redis.Client) *WeatherService {
	return &WeatherService{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (s *WeatherService) GetWeather(city string) (*CurrentWeatherResponse, error) {
	// Try to get from cache first if Redis is available
	if s.redisClient != nil {
		cacheKey := fmt.Sprintf("%s%s", cacheKeyPrefix, city)
		cachedData, err := s.redisClient.Get(s.ctx, cacheKey).Result()

		if err == nil {
			// Cache hit
			var weather CurrentWeatherResponse
			if err := json.Unmarshal([]byte(cachedData), &weather); err == nil {
				return &weather, nil
			}
		}
	}

	// Cache miss, error, or Redis not available - fetch from API
	weather, err := GetWeatherCurrentData(city)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}

	// Cache the result if Redis is available
	if s.redisClient != nil {
		if weatherData, err := json.Marshal(weather); err == nil {
			s.redisClient.Set(s.ctx, fmt.Sprintf("%s%s", cacheKeyPrefix, city), weatherData, cacheDuration)
		}
	}

	return weather, nil
}
