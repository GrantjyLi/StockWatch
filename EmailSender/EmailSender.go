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
	m.SetHeader("To", "1cbd49b946c6@maileroo-tester.com")
	m.SetHeader("Subject", "Your Alert has been triggered.")

	emailMessage := fmt.Sprintf("Alert %s was triggered.", alertData.alert_ID)

	m.SetBody("text/plain", emailMessage)

	d := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD)
	// d.SSL = true // enable for port 465

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email Sent" + alertData.alert_ID)
}
