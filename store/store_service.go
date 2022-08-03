package store

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

type DataStructure struct {
	Url        string `json:"url"`
	Counter    int32  `json:"counter"`
	Created_at string `json:"created_at"`
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = (24 * 7) * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "149.28.135.32:6379",
		Password: "4ran.net",
		DB:       12,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, created_at string) {
	saved := DataStructure{Url: originalUrl, Counter: 0, Created_at: created_at}

	savedJson, _ := json.Marshal(saved)
	err := storeService.redisClient.Set(shortUrl, savedJson, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func UpdateUrlMapping(shortUrl string, originalUrl string, created_at string, counter int32) {
	saved := DataStructure{Url: originalUrl, Counter: counter, Created_at: created_at}

	savedJson, _ := json.Marshal(saved)
	err := storeService.redisClient.Set(shortUrl, savedJson, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
