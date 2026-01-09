package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var (
	EMAIL_ADDR    string
	SMTP_USER     string
	SMTP_PASSWORD string
	SMTP_HOST     string
	SMTP_PORT     int
)

func setupSMTP() {
	EMAIL_ADDR = os.Getenv("EMAIL_ADDR")
	SMTP_USER = os.Getenv("SMTP_USER")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))

}

func sendEmail(alertData *Triggered_Alert) {

	m := gomail.NewMessage()
	m.SetHeader("From", EMAIL_ADDR)
	m.SetHeader("To", "ba105bca98c8@maileroo-tester.com") // testing
	// m.SetHeader("To", alertData.User_email)
	m.SetHeader("Subject", "Your Alert has been triggered.")

	emailMessage := fmt.Sprintf(
		"Alert %s was triggered.\n%s %s $%.2f",
		alertData.Alert_ID,
		alertData.Ticker,
		alertData.Operator,
		alertData.Target_price,
	)

	m.SetBody("text/plain", emailMessage)

	d := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD)
	// d.SSL = true // enable for port 465

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Email failed to send: %s\n", err.Error())
	}
	log.Printf("Email alert sent: %s to %s\n", alertData.Alert_ID, alertData.User_email)
}
