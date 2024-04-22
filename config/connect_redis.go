package config

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(config *ApiConfig) {
	opt, err := redis.ParseURL(config.RedisUrl)
	if err != nil {
		log.Fatal("Error parsing Redis URL: ", err)
		os.Exit(1)
	}
	config.RedisClient = redis.NewClient(opt)
}
