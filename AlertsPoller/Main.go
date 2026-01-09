package main

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var database *sql.DB
var FINNHUB_API_KEY string

const RMQ_RETRY_CONN_TIME = 5

func main() {

	database = DB_connect()
	defer database.Close()
	log.Println("Connected to database")

	allAlerts, _ := DB_getAlertTickers()
	log.Println("Alerts retreived")

	for {
		if RMQ_setup() {
			break
		}
		log.Println("RabbitMQ failed to connect")
		time.Sleep(RMQ_RETRY_CONN_TIME * time.Second)
	}
	log.Println("RabbitMQ conenction setup")

	FINNHUB_API_KEY = os.Getenv("FINNHUB_API_KEY")

	bootstrapPrices(allAlerts)
	log.Println("Initial prices retrieved")

	log.Println("Streaming current prices...")
	// getPriceUpdates(allAlerts)
	RMQ_close()
}
