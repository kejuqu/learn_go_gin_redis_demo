package config

import (
	"log"

	"github.com/go-redis/redis"

	"localhost/backend/global"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}


	global.RedisDB = RedisClient
}
