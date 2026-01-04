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

	// symbols := []string{
	// 	"AAPL",
	// 	"MSFT",
	// 	"AMZN",
	// 	"GOOG",
	// }

	FINNHUB_API_KEY = os.Getenv("FINNHUB_API_KEY")
	initRedis()

	bootstrapPrices(allAlerts)
	// getPriceUpdates(symbols)
}
