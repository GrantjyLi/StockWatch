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
	log.Println("Loaded environment variables")

	database = DB_connect()
	defer database.Close()
	log.Println("Connected to database")

	setupSMTP()
	log.Println("SMTP service setup")

	RMQ_setup()
	log.Println("RabbitMQ conenction setup")

	log.Println("Sending emails for incoming triggered alerts...")
	receiveNewAlert()
	RMQ_close()
}
