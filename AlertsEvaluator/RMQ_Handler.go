/*
Receives handling of rabitmq messages from the prices exchange
Evaluates if any users need to be alerted because of price updates
Sends necessary alerts to Email sender using the rabitmq alerts exchange
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	RMQ_PRICES_EX_NAME     = "new_prices"
	RMQ_PRICES_EX_KIND     = "topic"
	RMQ_PRICES_EX_DURABLE  = true
	RMQ_PRICES_EX_AUTODEL  = false
	RMQ_PRICES_EX_INTERNAL = false
	RMQ_PRICES_EX_NO_WAIT  = false
)

const (
	RMQ_ALERTS_EX_NAME     = "new_alerts"
	RMQ_ALERTS_EX_KIND     = "topic"
	RMQ_ALERTS_EX_TOPIC    = "alerts"
	RMQ_ALERTS_EX_DURABLE  = true
	RMQ_ALERTS_EX_AUTODEL  = false
	RMQ_ALERTS_EX_INTERNAL = false
	RMQ_ALERTS_EX_NO_WAIT  = false
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
	RMQ_CONN         *amqp.Connection
	RMQ_PRICES_CHANN *amqp.Channel
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
}

func setupPricesExchange() {
	err := RMQ_PRICES_CHANN.ExchangeDeclare(
		RMQ_PRICES_EX_NAME,
		RMQ_PRICES_EX_KIND,
		RMQ_PRICES_EX_DURABLE,
		RMQ_PRICES_EX_AUTODEL,
		RMQ_PRICES_EX_INTERNAL,
		RMQ_PRICES_EX_NO_WAIT,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	RMQ_QUEUE, err := RMQ_PRICES_CHANN.QueueDeclare(
		RMQ_QU_NAME,
		RMQ_QU_DURABLE,
		RMQ_QU_AUTODEL,
		RMQ_QU_INTERNAL,
		RMQ_QU_NO_WAIT,
		nil,
	)

	RMQ_PRICES_CHANN.QueueBind(
		RMQ_QUEUE.Name,
		RMQ_QU_ROUT_KEY,
		RMQ_PRICES_EX_NAME,
		RMQ_QU_NO_WAIT,
		nil,
	)

	RMQ_MSGS, _ = RMQ_PRICES_CHANN.Consume(
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

	RMQ_PRICES_CHANN, _ = RMQ_CONN.Channel()
	RMQ_ALERTS_CHANN, _ = RMQ_CONN.Channel()
	setupAlertsExchange()
	setupPricesExchange()

}

func RMQ_close() {
	if RMQ_PRICES_CHANN != nil {
		RMQ_PRICES_CHANN.Close()
	}
	if RMQ_ALERTS_CHANN != nil {
		RMQ_ALERTS_CHANN.Close()
	}
	if RMQ_CONN != nil {
		RMQ_CONN.Close()
	}
}

func publishNewAlert(newAlert *Triggered_Alert) error {
	alertData, _ := json.Marshal(newAlert)
	return RMQ_ALERTS_CHANN.Publish(
		RMQ_ALERTS_EX_NAME,
		(RMQ_ALERTS_EX_TOPIC + "." + newAlert.alert_ID),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        alertData,
		},
	)
}

func receiveNewPrice() {

	for msg := range RMQ_MSGS {
		var update PriceUpdate
		json.Unmarshal(msg.Body, &update)

		triggered_alerts, _ := DB_getAlertData(&update)
		for _, T := range triggered_alerts {
			publishNewAlert(T)
		}

		msg.Ack(false)
	}
}
