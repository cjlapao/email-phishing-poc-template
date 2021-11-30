package mailservice

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/cjlapao/common-go/executionctx"
)

var globalMailService *MailService

type SmtpMailServer struct {
	Host     string
	Port     int
	Username string
	Password string
}

type MailService struct {
	From   string
	Server *SmtpMailServer
}

func GetService() *MailService {
	if globalMailService != nil {
		return globalMailService
	}

	return nil
}

func NewMailService() *MailService {
	config := executionctx.GetConfigService()
	host := config.GetString("SMTP_HOST")
	port := config.GetInt("SMTP_PORT")
	userName := config.GetString("SMTP_USERNAME")
	password := config.GetString("SMTP_PASSWORD")
	globalMailService = &MailService{
		Server: &SmtpMailServer{host, port, userName, password},
	}

	return globalMailService
}

func (m *MailService) Send(to string, subject string, body string) {
	toArr := []string{to}
	auth := smtp.PlainAuth("", m.Server.Username, m.Server.Password, m.Server.Host)

	t, _ := template.New("Email").Parse(body)

	var bodyinBytes bytes.Buffer

	t.Execute(&bodyinBytes, struct {
		Url string
	}{
		Url: "http://localhost:10000/api/campain/",
	})

	err := smtp.SendMail(m.Server.Host+":"+string(m.Server.Port), auth, m.From, toArr, bodyinBytes.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("EmailSent")

}
