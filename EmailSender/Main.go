package main

import (
	"database/sql"
	"log"
	"time"
)

var database *sql.DB

func main() {

	database = DB_connect()
	defer database.Close()
	log.Println("Connected to database")

	setupSMTP()
	log.Println("SMTP service setup")

	for {
		if RMQ_setup() {
			break
		}
		log.Println("RabbitMQ failed to connect")
		time.Sleep(1 * time.Second)
	}
	log.Println("RabbitMQ conenction setup")

	log.Println("Sending emails for incoming triggered alerts...")
	receiveNewAlert()
	RMQ_close()
}
