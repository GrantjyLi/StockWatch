package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

const ENV_FILE = "AlertsEvaluator.env"

var database *sql.DB

func main() {

	err := godotenv.Load(ENV_FILE)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database = DB_connect()
	defer database.Close()

	RMQ_setup()
	receiveNewPrice()
	RMQ_close()
}
