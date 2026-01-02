package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

const ENV_FILE = "AlertsPoller.env"

var database *sql.DB

func main() {
	err := godotenv.Load(ENV_FILE)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database = DB_connect()
	defer database.Close()
	allAlerts, err := DB_getAlerts()

	if err != nil {
		fmt.Println("error getting alerts")
	}

	getTickerPrices(allAlerts)
}
