package main

import (
	"database/sql"
	"log"
	"time"
)

var database *sql.DB

const RMQ_RETRY_CONN_TIME = 5

func main() {

	database = DB_connect()
	defer database.Close()
	log.Println("Connected to database")

	for {
		if RMQ_setup() {
			break
		}
		log.Println("RabbitMQ failed to connect")
		time.Sleep(RMQ_RETRY_CONN_TIME * time.Second)
	}
	log.Println("RabbitMQ conenction setup")

	log.Println("Evaluating new prices...")
	receiveNewPrice()
	RMQ_close()
}
