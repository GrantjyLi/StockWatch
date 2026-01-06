package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

const ENV_FILE = "EmailSender.env"

func main() {
	err := godotenv.Load(ENV_FILE)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	EMAIL_ADDR := os.Getenv("EMAIL_ADDR")
	SMTP_USER := os.Getenv("SMTP_USER")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST := os.Getenv("SMTP_HOST")
	SMTP_PORT, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", EMAIL_ADDR, "StockWatch Alerts")
	m.SetHeader("To", "gnt.jy.li@gmail.com")
	m.SetHeader("Subject", "Test email")
	m.SetBody("text/plain", "This email was sent using Go.")

	d := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD)
	d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}
