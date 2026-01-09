package main

import (
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	RMQ_EX_NAME     = "new_prices"
	RMQ_EX_KIND     = "topic"
	RMQ_EX_TOPIC    = "stocks"
	RMQ_EX_DURABLE  = true
	RMQ_EX_AUTODEL  = false
	RMQ_EX_INTERNAL = false
	RMQ_EX_NO_WAIT  = false
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
		RMQ_EX_NAME,
		RMQ_EX_KIND,
		RMQ_EX_DURABLE,
		RMQ_EX_AUTODEL,
		RMQ_EX_INTERNAL,
		RMQ_EX_NO_WAIT,
		nil,
	)
	if err != nil {
		log.Printf("Failed to start/connect to %s: %s\n", RMQ_EX_NAME, err.Error())
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

func publishNewPrice(newUpdate *TickerData) error {
	tickerData, _ := json.Marshal(newUpdate)
	return RMQ_CHANN.Publish(
		RMQ_EX_NAME,
		(RMQ_EX_TOPIC + "." + newUpdate.Ticker),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        tickerData,
		},
	)
}
