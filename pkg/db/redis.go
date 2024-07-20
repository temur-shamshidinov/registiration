package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisCli() *redis.Client {
	option := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	redisCli := redis.NewClient(option)

	msg, err := redisCli.Ping(ctx).Result()
	if err != nil {
		log.Println("Error on connecting to Redis. Error:", err)
		return nil
	}

	log.Println("Successfully connected to Redis-cli, msg:", msg)
	return redisCli
}
