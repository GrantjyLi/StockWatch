package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
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
	Ticker string  `json:"s"`
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

func storePrice(symbol string, price float32) {
	data, _ := json.Marshal(PriceCache{
		Price: price,
		TS:    time.Now().Unix(),
	})

	err := rdb.Set(ctx, symbol, data, 0).Err()
	if err != nil {
		log.Println("Redis error:", err)
	}
}

/* REST init ============================ */

func fetchInitialPrice(alert *Alert) {
	ticker := alert.ticker
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

	if quote.Current > 0 {
		storePrice(ticker, quote.Current)
		publishNewPrice(&TickerData{
			Ticker: ticker,
			Price:  quote.Current,
		})
	}
}

/* Rate limit init ============================ */

func bootstrapPrices(alerts []*Alert) {
	var initPrices_WG sync.WaitGroup

	timeTicker := time.NewTicker(time.Second / FINNHUB_RATE_HZ)
	defer timeTicker.Stop()

	for _, alert := range alerts {
		<-timeTicker.C

		initPrices_WG.Add(1)
		go fetchInitialPrice(alert)

		go func(a *Alert) {
			defer initPrices_WG.Done()
			fetchInitialPrice(a)
		}(alert)
	}

	initPrices_WG.Wait()
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

		if msg.Type == "trade" {
			for _, trade := range msg.Data {
				storePrice(trade.Ticker, trade.Price)
				publishNewPrice(&trade)
			}
		}
	}
}
