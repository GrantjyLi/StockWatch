package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type PriceCache struct {
	Price float32 `json:"price"`
	TS    int64   `json:"ts"`
}

type QuoteResponse struct {
	Current float32 `json:"c"`
	High    float32 `json:"h"`
	Low     float32 `json:"l"`
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
	FINNHUB_QOUTE_EP = "https://finnhub.io/api/v1/quote"
	FINNHUB_WS       = "wss://ws.finnhub.io?token="
	REDIS_ADDR       = "localhost:6379"
	FINNHUB_RATE_HZ  = 30
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

func storePrice(symbol string, qoute QuoteResponse) {
	data, _ := json.Marshal(PriceCache{
		Price: qoute.Current,
		TS:    time.Now().Unix(),
	})

	err := rdb.Set(ctx, symbol, data, 0).Err()
	if err != nil {
		log.Println("Redis error:", err)
	}

	val, err := rdb.Get(ctx, symbol).Result()
	if err != nil {
		log.Println("Redis error:", err)
	}
	fmt.Printf("Value from redis: %s\n", val)
}

/* REST init ============================ */

func fetchInitialPrice(ticker string) {
	url := fmt.Sprintf("%s?symbol=%s&token=%s", FINNHUB_QOUTE_EP, ticker, FINNHUB_API_KEY)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Decode error:", err)
		return
	}
	defer resp.Body.Close()

	var quote QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		log.Println("Decode error:", err)
		return
	}

	fmt.Printf("%s: $%.2f\n", ticker, quote.Current)
	if quote.Current > 0 {
		storePrice(ticker, quote)
	}
}

/* Rate limit init ============================ */

func bootstrapPrices(alerts []*Alert) {
	timeTicker := time.NewTicker(time.Second / FINNHUB_RATE_HZ)
	defer timeTicker.Stop()

	for _, alert := range alerts {
		<-timeTicker.C
		fetchInitialPrice(alert.ticker)
	}
}

/* web socket streaming ============================ */

func getPriceUpdates(symbols []string) {

	ws, _, err := websocket.DefaultDialer.Dial(FINNHUB_WS+FINNHUB_API_KEY, nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		ws.WriteMessage(websocket.TextMessage, msg)
	}
	for {
		var msg TradeMsg
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Fatal("WS read error:", err)
		}
		fmt.Println("Message from server ", msg)

		// if msg.Type == "trade" {
		// 	for _, trade := range msg.Data {
		// 		storePrice(trade.Symbol, trade.Price)
		// 	}
		// }
	}
}
