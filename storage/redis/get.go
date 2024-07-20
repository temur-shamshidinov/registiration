package redis

import (
	"log"
	"registiration/pkg/db"
)

// Get - Redisdan ma'lumotni oladi

func Get(key string) (string, error) {

	redisCli := db.NewRedisCli()

	data, err := redisCli.Get(ctx, key).Result()
	if err != nil {
		log.Println("error in get  data to redis!", err)
		return data, nil
	}

	log.Printf("Successfully got value for key %s: %s", key, data)
	return data, nil
}
