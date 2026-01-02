package main

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

const (
	finnhubToken = os.Getenv("FINNHUB_API_KEY")
	finnhubREST  = "https://finnhub.io/api/v1/quote"
	finnhubWS    = "wss://ws.finnhub.io?token="
	REDIS_ADDR   = "localhost:6379"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: REDIS_ADDR,
	})
}
