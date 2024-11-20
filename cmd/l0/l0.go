package main

import (
	"ru/zakat/L0/cache"
	"ru/zakat/L0/db"
	"ru/zakat/L0/internal/router"
	"ru/zakat/L0/kafka"
	"ru/zakat/L0/logger"
)

func main() {
	log := logger.NewLogger()
	repo := db.NewRepository()
	cache := cache.NewCache(repo)
	orderProducer := kafka.NewProducer()

	go kafka.NewConsumer(cache)

	router.InitRouter(log, cache, orderProducer)
}
