package mailservice

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strconv"

	"github.com/cjlapao/common-go/executionctx"
)

var globalMailService *MailService

type SmtpMailServer struct {
	Host     string
	Port     int
	Username string
	Password string
	TLS      bool
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
	tls := config.GetBool("SMTP_TLS_ENABLED")
	globalMailService = &MailService{
		Server: &SmtpMailServer{host, port, userName, password, tls},
	}

	return globalMailService
}

func (m *MailService) Send(toAddress string, subject string, body string) {
	from := mail.Address{"", "cjlapao@gmail.com"}
	to := mail.Address{"", toAddress}
	subj := "This is the email subject"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	serverHost := m.Server.Host + ":" + strconv.Itoa(m.Server.Port)
	auth := smtp.PlainAuth("", m.Server.Username, m.Server.Password, m.Server.Host)
	err := smtp.SendMail(serverHost, auth, headers["from"], []string{toAddress}, []byte(message))
	if err != nil {
		log.Panic(err)
	}
}
