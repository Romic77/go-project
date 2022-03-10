package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	_, err = redisDB.Ping().Result()
	return
}
func main() {
	err := initRedis()
	if err != nil {
		return
	}

	redisDB.Set("score", 100, 0).Err()
	result, _ := redisDB.Get("score").Result()
	fmt.Printf("%v", result)
}
