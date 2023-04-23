package redis

import "github.com/redis/go-redis/v9"

var Cache *redis.Client

func SetUpRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "cache_service:6379",
		DB:   0,
	})
}
