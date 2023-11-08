package services

import (
	"gopkg.in/gomail.v2"
	"log"
	"meetme/be/config"
)

type Mailer struct{}

func (m *Mailer) Send(message *gomail.Message) {
	message.SetHeader("From", "etracking.th@gmail.com")

	if err := config.Mailer.DialAndSend(message); err != nil {
		log.Panicln("[Mailer] ", err)
	}
}
