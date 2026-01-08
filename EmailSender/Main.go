package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

const ENV_FILE = "EmailSender.env"

var database *sql.DB

func main() {
	err := godotenv.Load(ENV_FILE)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database = DB_connect()
	defer database.Close()

	setupSMTP()

	RMQ_setup()
	receiveNewAlert()
	RMQ_close()
}
