package config

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

// Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	auth := smtp.PlainAuth("", os.Getenv("MAILER_USERNAME"), os.Getenv("MAILER_PASSWORD"), os.Getenv("MAILER_HOST"))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from := fmt.Sprintf("From: meetme play <%s>\r\n", os.Getenv("MAILER_USERNAME"))
	to := fmt.Sprintf("To: %s\r\n", strings.Join(r.to, " "))
	subject := fmt.Sprintf("Subject: %s!\n", r.subject)
	msg := []byte(subject + from + to + mime + "\n" + r.body)
	addr := os.Getenv("MAILER_HOST") + ":" + os.Getenv("MAILER_PORT")

	if err := smtp.SendMail(addr, auth, os.Getenv("MAILER_USERNAME"), r.to, msg); err != nil {
		log.Print(err)
		return false, err
	}
	return true, nil

}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
