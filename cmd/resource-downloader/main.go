package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"

	"golang2/internal/infrastructure/external"

	internalRedis "golang2/internal/infrastructure/redis"
)

func main() {
	fmt.Println("Repository pattern example")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "secret",
		DB:       0,
	})
	cache := internalRedis.NewRedisDownloaderRepository(redisClient)
	userServiceClient := external.NewDefaultUserServiceClient("http://localhost")
	external.NewCachedExternalDownloaderRepository(
		&userServiceClient, cache,
	)
}
