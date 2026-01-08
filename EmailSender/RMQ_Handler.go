package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Triggered_Alert struct {
	alert_ID   string
	user_email string
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
	RMQ_QU_NAME     = "alerts_email"
	RMQ_QU_ROUT_KEY = "alerts.#"
	RMQ_QU_DURABLE  = true
	RMQ_QU_AUTODEL  = false
	RMQ_QU_INTERNAL = false
	RMQ_QU_NO_WAIT  = false
)

var (
	RMQ_CONN         *amqp.Connection
	RMQ_ALERTS_CHANN *amqp.Channel
	RMQ_MSGS         <-chan amqp.Delivery
)

func setupAlertsExchange() {
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
		log.Fatal(err)
	}

	RMQ_QUEUE, err := RMQ_ALERTS_CHANN.QueueDeclare(
		RMQ_QU_NAME,
		RMQ_QU_DURABLE,
		RMQ_QU_AUTODEL,
		RMQ_QU_INTERNAL,
		RMQ_QU_NO_WAIT,
		nil,
	)

	RMQ_ALERTS_CHANN.QueueBind(
		RMQ_QUEUE.Name,
		RMQ_QU_ROUT_KEY,
		RMQ_ALERTS_EX_NAME,
		RMQ_QU_NO_WAIT,
		nil,
	)

	RMQ_MSGS, _ = RMQ_ALERTS_CHANN.Consume(
		RMQ_QUEUE.Name,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)
}

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

	RMQ_ALERTS_CHANN, _ = RMQ_CONN.Channel()
	setupAlertsExchange()
}

func RMQ_close() {
	if RMQ_ALERTS_CHANN != nil {
		RMQ_ALERTS_CHANN.Close()
	}
	if RMQ_CONN != nil {
		RMQ_CONN.Close()
	}
}

func receiveNewAlert() {

	for msg := range RMQ_MSGS {
		var update Triggered_Alert
		json.Unmarshal(msg.Body, &update)

		fmt.Println("Received alert" + update.alert_ID)

		go sendEmail(&update)

		msg.Ack(false)
	}
}
