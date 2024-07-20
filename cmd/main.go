package main

import (
	"log"
	"registiration/api"
	"registiration/config"
	"registiration/pkg/db"
	"registiration/storage"
)

func main() {
	cfg := config.Load()

	dbp, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Println("error on connect to ConToDb:", err)
		return
	}

	redisCli := db.NewRedisCli()

	if redisCli == nil {
		log.Println("error to create Redis client")
		return
	}

	storage := storage.NewStorage(dbp)

	api.Api(storage)

}
