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
	log.Println("Loaded environment variables")

	database = DB_connect()
	defer database.Close()
	log.Println("Connected to database")

	RMQ_setup()
	log.Println("RabbitMQ conenction setup")

	log.Println("Evaluating new prices...")
	receiveNewPrice()
	RMQ_close()
}
