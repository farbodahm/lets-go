package database

import (
	"log"

	"github.com/go-redis/redis"
)

const (
	REDIS_ADDR = "redis"
	REDIS_PORT = "6379"
	REDIS_PASS = ""
	REDIS_DB   = 0
)

// Main connection to interact with Redis database
var RedisClient *redis.Client

// Initialize first connection to Redis database
func InitializeRedisConnection() {
	// Create connection
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR + ":" + REDIS_PORT,
		Password: REDIS_PASS,
		DB:       REDIS_DB,
	})

	// Validate Connection
	if RedisClient == nil {
		log.Fatal("Error: Couldn't connect to Redis server")
	}
	result, err := RedisClient.Ping().Result()
	if err != nil || result != "PONG" {
		log.Fatal("Error: Couldn't ping Redis server: ", err)
	}
}
