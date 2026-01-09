package main

import (
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	RMQ_WLISTS_EX_NAME     = "new_watchlists"
	RMQ_WLISTS_EX_KIND     = "topic"
	RMQ_WLISTS_EX_TOPIC    = "watchlist"
	RMQ_WLISTS_EX_DURABLE  = true
	RMQ_WLISTS_EX_AUTODEL  = false
	RMQ_WLISTS_EX_INTERNAL = false
	RMQ_WLISTS_EX_NO_WAIT  = false
)

var (
	RMQ_CONN  *amqp.Connection
	RMQ_CHANN *amqp.Channel
)

func RMQ_setup() bool {
	var err error
	RMQ_CONN, err = amqp.Dial(os.Getenv("RMQ_ADDR_URL"))
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err.Error())
		return false
	}

	RMQ_CHANN, _ = RMQ_CONN.Channel()

	err = RMQ_CHANN.ExchangeDeclare(
		RMQ_WLISTS_EX_NAME,
		RMQ_WLISTS_EX_KIND,
		RMQ_WLISTS_EX_DURABLE,
		RMQ_WLISTS_EX_AUTODEL,
		RMQ_WLISTS_EX_INTERNAL,
		RMQ_WLISTS_EX_NO_WAIT,
		nil,
	)
	if err != nil {
		log.Printf("Failed to start/connect to %s: %s\n", RMQ_WLISTS_EX_NAME, err.Error())
		return false
	}

	return true
}

func RMQ_close() {
	if RMQ_CHANN != nil {
		RMQ_CHANN.Close()
	}
	if RMQ_CONN != nil {
		RMQ_CONN.Close()
	}
}

func publishNewWatchlist(newWatchlistReq *CreateWatchlistRequest_t) error {
	watchlistReqData, _ := json.Marshal(newWatchlistReq)
	return RMQ_CHANN.Publish(
		RMQ_WLISTS_EX_NAME,
		(RMQ_WLISTS_EX_TOPIC + "." + newWatchlistReq.WatchlistData.ID),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        watchlistReqData,
		},
	)
}
