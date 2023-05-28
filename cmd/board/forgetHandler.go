package main

import (
	"fmt"
	"os"
    "net/smtp"
	"net/http"
    "strings"

	"gopkg.in/yaml.v2"
)

type MailSettings struct {
	Hostname string `yaml:"mail_host"`
	MailAddress string `yaml:"mail_address"`
	MailPassword string `yaml:"mail_password"`
}

func forgetHandler(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		panic("error: parse error occured.")
	}

	email := r.Form.Get("email")

	mailSettings := MailSettings{}
	b, _ := os.ReadFile("mail.yaml")
	yaml.Unmarshal(b, &mailSettings)

    from := mailSettings.MailAddress
    recipients := []string{email}
    subject := "Mail Test"
    body := "Hello World!\nHello Gopher!"

    auth := smtp.CRAMMD5Auth(mailSettings.MailAddress, mailSettings.MailPassword)
    msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(recipients, ","), subject, body), "\n", "\r\n"))

	err := smtp.SendMail(fmt.Sprintf("%s:%d", mailSettings.Hostname, 587), auth, from, recipients, msg);
    if err != nil {
        panic("error : failed to send mail.")
    }
	fmt.Fprintln(w, "{\"success\": \"true\"}")
}
