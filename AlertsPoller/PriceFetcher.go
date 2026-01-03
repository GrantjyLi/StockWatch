package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

type PriceCache struct {
	Price float32 `json:"price"`
	TS    int64   `json:"ts"`
}

type QuoteResponse struct {
	Current float32 `json:"c"`
}

type TradeMsg struct {
	Type string       `json:"type"`
	Data []TickerData `json:"data"`
}

type TickerData struct {
	Symbol string  `json:"s"`
	Price  float32 `json:"p"`
}

const (
	finnhubREST     = "https://finnhub.io/api/v1/quote"
	FINNHUB_WS      = "wss://ws.finnhub.io?token="
	REDIS_ADDR      = "localhost:6379"
	FINNHUB_RATE_HZ = 30
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

/* Price Caching ============================ */

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: REDIS_ADDR,
	})
}

func storePrice(symbol string, price float32) {
	data, _ := json.Marshal(PriceCache{
		Price: price,
		TS:    time.Now().Unix(),
	})

	err := rdb.Set(ctx, "price:"+symbol, data, 0).Err()
	if err != nil {
		log.Println("Redis error:", err)
	}
}

/* REST init ============================ */

func fetchInitialPrice(symbol string) {
	req, _ := http.NewRequest("GET", finnhubREST, nil)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("token", FINNHUB_API_KEY)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("REST error:", err)
		return
	}
	defer resp.Body.Close()

	var quote QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		log.Println("Decode error:", err)
		return
	}

	if quote.Current > 0 {
		storePrice(symbol, quote.Current)
	}
}

/* Rate limit init ============================ */

func bootstrapPrices(symbols []string) {
	ticker := time.NewTicker(time.Second / FINNHUB_RATE_HZ)
	defer ticker.Stop()

	for _, symbol := range symbols {
		<-ticker.C
		go fetchInitialPrice(symbol)
	}
}

/* web socket streaming ============================ */

func startWebSocket(symbols []string) {
	fmt.Println(FINNHUB_WS + FINNHUB_API_KEY)
	ws, _, err := websocket.DefaultDialer.Dial(
		FINNHUB_WS+FINNHUB_API_KEY, nil,
	)
	if err != nil {
		log.Fatal("WS error:", err)
	}
	defer ws.Close()

	// Subscribe
	for _, s := range symbols {
		msg := map[string]string{
			"type":   "subscribe",
			"symbol": s,
		}
		ws.WriteJSON(msg)
	}

	for {
		var msg TradeMsg
		if err := ws.ReadJSON(&msg); err != nil {
			log.Fatal("WS read error:", err)
		}

		if msg.Type == "trade" {
			for _, trade := range msg.Data {
				storePrice(trade.Symbol, trade.Price)
			}
		}
	}
}
