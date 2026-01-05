package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PriceUpdate struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

func receiveNewPrice() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(
		"alerts_eval",
		true,
		false,
		false,
		false,
		nil,
	)

	ch.QueueBind(
		q.Name,
		"", // empty = all routing keys
		"prices",
		false,
		nil,
	)

	msgs, _ := ch.Consume(
		q.Name,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)

	log.Println("Waiting for price updates...")

	for msg := range msgs {
		var update PriceUpdate
		json.Unmarshal(msg.Body, &update)

		log.Printf("Evaluating alerts for %s @ %.2f\n", update.Ticker, update.Price)

		// TODO:
		// alerts := loadAlerts(update.Ticker)
		// evaluate(alerts, update.Price)

		msg.Ack(false)
	}
}
