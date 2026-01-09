package main

import (
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Triggered_Alert struct {
	Alert_ID     string  `json:"alert_ID"`
	Ticker       string  `json:"ticker"`
	Target_price float32 `json:"target_price"`
	Operator     string  `json:"operator"`
	User_email   string  `json:"user_email"`
}
type CreateWatchlistRequest_t struct {
	UserID        string    `json:"userID"`
	User_email    string    `json:"email"`
	WatchlistData Watchlist `json:"watchlistData"`
}
type Watchlist struct {
	ID     string   `json:"ID"`
	Name   string   `json:"name"`
	Alerts []*Alert `json:"alerts"`
}
type Alert struct {
	ID       string  `json:"ID"`
	Ticker   string  `json:"ticker"`
	Operator string  `json:"operator"`
	Price    float32 `json:"price"`
}

const (
	RMQ_ALERTS_EX_NAME     = "new_alerts"
	RMQ_ALERTS_EX_KIND     = "topic"
	RMQ_ALERTS_EX_DURABLE  = true
	RMQ_ALERTS_EX_AUTODEL  = false
	RMQ_ALERTS_EX_INTERNAL = false
	RMQ_ALERTS_EX_NO_WAIT  = false
)

const (
	RMQ_ALERTS_QU_NAME     = "alerts_email"
	RMQ_ALERTS_QU_ROUT_KEY = "alerts.#"
	RMQ_ALERTS_QU_DURABLE  = true
	RMQ_ALERTS_QU_AUTODEL  = false
	RMQ_ALERTS_QU_INTERNAL = false
	RMQ_ALERTS_QU_NO_WAIT  = false
)

const (
	RMQ_WLISTS_EX_NAME     = "new_watchlists"
	RMQ_WLISTS_EX_KIND     = "topic"
	RMQ_WLISTS_EX_DURABLE  = true
	RMQ_WLISTS_EX_AUTODEL  = false
	RMQ_WLISTS_EX_INTERNAL = false
	RMQ_WLISTS_EX_NO_WAIT  = false
)

const (
	RMQ_WLISTS_QU_NAME     = "watchlist_email"
	RMQ_WLISTS_QU_ROUT_KEY = "watchlist.#"
	RMQ_WLISTS_QU_DURABLE  = true
	RMQ_WLISTS_QU_AUTODEL  = false
	RMQ_WLISTS_QU_INTERNAL = false
	RMQ_WLISTS_QU_NO_WAIT  = false
)

var (
	RMQ_CONN         *amqp.Connection
	RMQ_ALERTS_CHANN *amqp.Channel
	RMQ_WLISTS_CHANN *amqp.Channel
	RMQ_ALERTS_MSGS  <-chan amqp.Delivery
	RMQ_WLISTS_MSGS  <-chan amqp.Delivery
)

func setupAlertsExchange() bool {
	err := RMQ_ALERTS_CHANN.ExchangeDeclare(
		RMQ_ALERTS_EX_NAME,
		RMQ_ALERTS_EX_KIND,
		RMQ_ALERTS_EX_DURABLE,
		RMQ_ALERTS_EX_AUTODEL,
		RMQ_ALERTS_EX_INTERNAL,
		RMQ_ALERTS_EX_NO_WAIT,
		nil,
	)
	if err != nil {
		log.Printf("Failed to start/connect to %s: %s", RMQ_ALERTS_EX_NAME, err.Error())
		return false
	}
	log.Println("Connected to exchange: ", RMQ_ALERTS_EX_NAME)

	RMQ_ALERTS_QUEUE, err := RMQ_ALERTS_CHANN.QueueDeclare(
		RMQ_ALERTS_QU_NAME,
		RMQ_ALERTS_QU_DURABLE,
		RMQ_ALERTS_QU_AUTODEL,
		RMQ_ALERTS_QU_INTERNAL,
		RMQ_ALERTS_QU_NO_WAIT,
		nil,
	)

	if err != nil {
		log.Printf("Failed to start/connect to %s: %s", RMQ_ALERTS_QU_NAME, err.Error())
		return false
	}

	err = RMQ_ALERTS_CHANN.QueueBind(
		RMQ_ALERTS_QUEUE.Name,
		RMQ_ALERTS_QU_ROUT_KEY,
		RMQ_ALERTS_EX_NAME,
		RMQ_ALERTS_QU_NO_WAIT,
		nil,
	)

	if err != nil {
		log.Printf("Failed to start/connect to queue %s: %s", RMQ_ALERTS_QUEUE.Name, err.Error())
		return false
	}
	log.Println("Fully bond to queue: ", RMQ_ALERTS_QU_NAME)

	RMQ_ALERTS_MSGS, _ = RMQ_ALERTS_CHANN.Consume(
		RMQ_ALERTS_QUEUE.Name,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)

	return true
}

func setupWatchlistsExchange() bool {
	err := RMQ_WLISTS_CHANN.ExchangeDeclare(
		RMQ_WLISTS_EX_NAME,
		RMQ_WLISTS_EX_KIND,
		RMQ_WLISTS_EX_DURABLE,
		RMQ_WLISTS_EX_AUTODEL,
		RMQ_WLISTS_EX_INTERNAL,
		RMQ_WLISTS_EX_NO_WAIT,
		nil,
	)
	if err != nil {
		log.Printf("Failed to start/connect to %s: %s", RMQ_WLISTS_EX_NAME, err.Error())
		return false
	}
	log.Println("Connected to exchange: ", RMQ_WLISTS_EX_NAME)

	RMQ_WLISTS_QUEUE, err := RMQ_WLISTS_CHANN.QueueDeclare(
		RMQ_WLISTS_QU_NAME,
		RMQ_WLISTS_QU_DURABLE,
		RMQ_WLISTS_QU_AUTODEL,
		RMQ_WLISTS_QU_INTERNAL,
		RMQ_WLISTS_QU_NO_WAIT,
		nil,
	)

	if err != nil {
		log.Printf("Failed to start/connect to %s: %s", RMQ_WLISTS_QU_NAME, err.Error())
		return false
	}

	err = RMQ_WLISTS_CHANN.QueueBind(
		RMQ_WLISTS_QUEUE.Name,
		RMQ_WLISTS_QU_ROUT_KEY,
		RMQ_WLISTS_EX_NAME,
		RMQ_WLISTS_QU_NO_WAIT,
		nil,
	)

	if err != nil {
		log.Printf("Failed to start/connect to queue %s: %s", RMQ_WLISTS_QUEUE.Name, err.Error())
		return false
	}
	log.Println("Fully bond to queue: ", RMQ_WLISTS_QU_NAME)

	RMQ_WLISTS_MSGS, _ = RMQ_WLISTS_CHANN.Consume(
		RMQ_WLISTS_QUEUE.Name,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)

	return true
}

func RMQ_setup() bool {
	var err error
	RMQ_CONN, err = amqp.Dial(os.Getenv("RMQ_ADDR_URL"))
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err.Error())
		return false
	}

	RMQ_ALERTS_CHANN, _ = RMQ_CONN.Channel()
	RMQ_WLISTS_CHANN, _ = RMQ_CONN.Channel()

	if !setupAlertsExchange() {
		return false
	}
	if !setupWatchlistsExchange() {
		return false
	}

	return true
}

func RMQ_close() {
	if RMQ_ALERTS_CHANN != nil {
		RMQ_ALERTS_CHANN.Close()
	}
	if RMQ_WLISTS_CHANN != nil {
		RMQ_WLISTS_CHANN.Close()
	}
	if RMQ_CONN != nil {
		RMQ_CONN.Close()
	}
}

func receiveNewAlert() {

	for msg := range RMQ_ALERTS_MSGS {
		var update Triggered_Alert
		json.Unmarshal(msg.Body, &update)

		go func() {
			sendAlertEmail(&update)
			// DB_AlertTriggered(&update)
		}()

		msg.Ack(false)
	}
}

func receiveNewWatchlist() {
	for msg := range RMQ_WLISTS_MSGS {
		var update CreateWatchlistRequest_t
		json.Unmarshal(msg.Body, &update)

		go func() {
			sendWatchlistEmail(&update)
		}()

		msg.Ack(false)
	}
}
