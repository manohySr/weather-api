package cache

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()
var addr = os.Getenv("addr")

func InitRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		return nil
	}

	log.Println("connected to Redis!")
	return rdb
}
