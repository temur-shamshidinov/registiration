package redis

import (
	"context"
	"log"
	"registiration/pkg/db"

	"time"
)

var ctx = context.Background()

// set - redisga ma`lumotni saqlaydi

func Set(key string, value string) error {

	redisCli := db.NewRedisCli()

	err := redisCli.Set(ctx, key, value, 1*time.Minute).Err()
	if err != nil {
		log.Printf("error to set value for key %s. error: %v", key, err)
		return err
	}

	log.Printf("Successfully set value for key %s", key)
	return nil
}
