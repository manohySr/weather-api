# Weather API Wrapper Service

A weather API wrapper service that fetches weather data from Visual Crossing API with Redis caching support. This project demonstrates working with third-party APIs, implementing caching, and handling environment variables.

## Features

- 🌤️ Real-time weather data from Visual Crossing API
- 💾 Redis caching for improved performance
- 🚀 RESTful API endpoints
- 🔒 Environment variable configuration
- 🐳 Docker support for Redis

![image](https://github.com/user-attachments/assets/78f04cab-118f-457f-b52c-c5c87699361a)


## Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Visual Crossing API key

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
key=your_visual_crossing_api_key
addr=localhost:6379
PORT=3000
```

## Project Structure
```
https://gitingest.com/manohySr/weather-api
```

```
.
├── cache/
│   └── redis.go        # Redis client configuration
├── weather/
│   ├── client.go       # Visual Crossing API client
│   └── service.go      # Weather service with caching
├── main.go             # API server and routes
├── docker-compose.yml  # Redis container configuration
└── README.md
```

## API Endpoints

### Get Weather by City

```http
GET /weather/:city
```

#### Response

```json
{
    "data": {
        "datetime": "14:17:00",
        "temp": 66,
        "feelslike": 66,
        "humidity": 60,
        "precip": 0,
        "precipprob": 0,
        "windgust": 12.2,
        "windspeed": 6.1,
        "winddir": 259,
        "pressure": 1010,
        "visibility": 6.2,
        "cloudcover": 88,
        "uvindex": 1,
        "conditions": "Partially cloudy"
    },
    "city": "london"
}
```

#### Status Codes

- `200 OK`: Weather data retrieved successfully
- `404 Not Found`: City not found
- `500 Internal Server Error`: Server error or API failure

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd weather-api
```

2. Start Redis using Docker Compose:
```bash
docker-compose up -d
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your API key
```

4. Run the application:
```bash
go run main.go
```

The server will start at `http://localhost:3000`

## Caching

The service implements Redis caching with the following features:
- 30-minute cache duration for weather data
- Automatic cache invalidation
- Fallback to API calls when Redis is unavailable

## Error Handling

The service includes comprehensive error handling:
- Invalid city names
- API failures
- Redis connection issues
- Proper HTTP status codes and error messages

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Visual Crossing Weather API](https://www.visualcrossing.com/)
- [Go Fiber](https://gofiber.io/)
- [Redis](https://redis.io/)
- [roadmap.sh Weather API Project](https://roadmap.sh/projects/weather-api-wrapper-service) - Project inspiration and requirements 
