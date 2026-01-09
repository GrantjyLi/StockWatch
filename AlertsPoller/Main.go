package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// todo: investigate multiple queries for 1 ticker
const ENV_FILE = "AlertsPoller.env"

var database *sql.DB
var FINNHUB_API_KEY string

func main() {
	err := godotenv.Load(ENV_FILE)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded environment variables")

	database = DB_connect()
	defer database.Close()
	log.Println("Connected to database")

	allAlerts, err := DB_getAlertTickers()
	log.Println("Alerts retreived")

	RMQ_setup()
	log.Println("RabbitMQ conenction setup")

	FINNHUB_API_KEY = os.Getenv("FINNHUB_API_KEY")

	bootstrapPrices(allAlerts)
	log.Println("Initial prices retrieved")

	log.Println("Streaming current prices...")
	getPriceUpdates(allAlerts)
	RMQ_close()
}
