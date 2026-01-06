package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PriceUpdate struct {
	Ticker string  `json:"s"`
	Price  float32 `json:"p"`
}

const (
	RMQ_EX_NAME     = "new_prices"
	RMQ_EX_KIND     = "topic"
	RMQ_EX_DURABLE  = true
	RMQ_EX_AUTODEL  = false
	RMQ_EX_INTERNAL = false
	RMQ_EX_NO_WAIT  = false
)

const (
	RMQ_QU_NAME     = "alerts_evaluator"
	RMQ_QU_ROUT_KEY = "stocks.#"
	RMQ_QU_DURABLE  = true
	RMQ_QU_AUTODEL  = false
	RMQ_QU_INTERNAL = false
	RMQ_QU_NO_WAIT  = false
)

var (
	RMQ_CONN  *amqp.Connection
	RMQ_CHANN *amqp.Channel
	RMQ_MSGS  <-chan amqp.Delivery
)

func RMQ_setup() {
	RMQ_address := fmt.Sprintf(
		"amqp://%s:%s@%s/",
		os.Getenv("RMQ_UN"),
		os.Getenv("RMQ_PW"),
		os.Getenv("RMQ_ADDR"),
	)
	var err error
	RMQ_CONN, err = amqp.Dial(RMQ_address)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	RMQ_QUEUE, err := RMQ_CHANN.QueueDeclare(
		RMQ_QU_NAME,
		RMQ_QU_DURABLE,
		RMQ_QU_AUTODEL,
		RMQ_QU_INTERNAL,
		RMQ_QU_NO_WAIT,
		nil,
	)

	RMQ_CHANN.QueueBind(
		RMQ_QUEUE.Name,
		RMQ_QU_ROUT_KEY,
		RMQ_EX_NAME,
		RMQ_QU_NO_WAIT,
		nil,
	)

	RMQ_MSGS, _ = RMQ_CHANN.Consume(
		RMQ_QUEUE.Name,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)
}

func RMQ_close() {
	if RMQ_CHANN != nil {
		RMQ_CHANN.Close()
	}
	if RMQ_CONN != nil {
		RMQ_CONN.Close()
	}
}

func receiveNewPrice() {

	for msg := range RMQ_MSGS {
		var update PriceUpdate
		json.Unmarshal(msg.Body, &update)

		log.Printf("Evaluating alerts for %s @ %.2f\n", update.Ticker, update.Price)

		triggered_alerts, _ := DB_getAlertData(&update)
		for _, T := range triggered_alerts {
			fmt.Printf("Triggered alert %s for user %s\n", T.alert_ID, T.user_email)
		}

		msg.Ack(false)
	}
}
