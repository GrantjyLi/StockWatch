package main

import (
	"database/sql"
	// "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const ENV_FILE = "AlertsPoller.env"

var database *sql.DB
var FINNHUB_API_KEY string

func main() {
	err := godotenv.Load(ENV_FILE)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database = DB_connect()
	defer database.Close()
	allAlerts, err := DB_getAlerts()

	RMQ_setup()

	FINNHUB_API_KEY = os.Getenv("FINNHUB_API_KEY")
	initRedis()

	bootstrapPrices(allAlerts)
	// getPriceUpdates(symbols)
	RMQ_close()
}
