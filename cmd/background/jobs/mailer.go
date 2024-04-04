package jobs

import (
	"air-e-notification-adviser/config"
	"air-e-notification-adviser/internal/caribesol/dto"
	"air-e-notification-adviser/internal/platform/files"
	"bytes"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"sync"
	"text/template"
)

type ContentType string

const TextPlainContentType = "text/plain"

type Mail struct {
	Host        string
	Port        int
	Username    string
	Password    string
	FromAddress string
	Wait        *sync.WaitGroup
	MailerChan  chan Message
	ErrorChan   chan error
	DoneChan    chan bool
}

func NewMail(cfg *config.EnvVars, wg *sync.WaitGroup) *Mail {
	return &Mail{
		Host:        cfg.SMTPHost,
		Port:        cfg.SMTPPort,
		Username:    cfg.SMTPHost,
		Password:    cfg.SMTPHost,
		FromAddress: cfg.SMTPHost,
		Wait:        wg,
		MailerChan:  make(chan Message, 100),
		ErrorChan:   make(chan error),
		DoneChan:    make(chan bool),
	}
}

type Message struct {
	From    string
	To      string
	Subject string
	Data    struct {
		ConsultarNICDTOResponse dto.ConsultarNICDTOResponse
		Datetime                string
	}
}

// listenForMail is a function to listen for messages on the MailerChan.
func (w *Worker) listenForMail(logger *zap.Logger) {
	for {
		select {
		case message := <-w.Mailer.MailerChan:
			go w.Mailer.sendMail(message, w.Mailer.ErrorChan)
		case err := <-w.Mailer.ErrorChan:
			logger.Error(err.Error())
		case <-w.Mailer.DoneChan:
			w.Mailer.Wait.Done()
			return
		}
	}
}

func (w *Worker) sendEmail(message Message) {
	w.Wait.Add(1)
	w.Mailer.MailerChan <- message
}

func (m *Mail) sendMail(message Message, errorChan chan error) {
	defer m.Wait.Done()

	if message.From == "" {
		message.From = m.FromAddress
	}

	plainTextMessage, err := m.buildPlainTextMessage(message)
	if err != nil {
		errorChan <- err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", message.From)
	mail.SetHeader("To", message.To)
	mail.SetHeader("Subject", message.Subject)
	mail.SetBody(TextPlainContentType, plainTextMessage)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	if err := dialer.DialAndSend(mail); err != nil {
		errorChan <- err
	}
}

func (m *Mail) buildPlainTextMessage(message Message) (string, error) {
	notificationTemplateFile, err := files.GetFile("templates/notification.plain")
	if err != nil {
		return "", err
	}

	t, err := template.New("email-plain").ParseFiles(notificationTemplateFile)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "body", message.Data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
