package gmail

import (
	"fmt"
	"net/smtp"
)

const (
	mailServerHost = "smtp.gmail.com"
	mailServerPort = 587
)

// Client holds information about gmail client
type Client struct {
	SenderGmailAddress string
	SenderPassword     string
	mailServer         string
}

// Message describes an email message which can be send
type Message struct {
	Receiver string
	Subject  string
	Body     string
}

// NewClient returns a client which can be used to send emails
func NewClient(username, password string) Client {
	c := Client{
		SenderGmailAddress: username,
		SenderPassword:     password,
		mailServer:         fmt.Sprintf("%s:%d", mailServerHost, mailServerPort),
	}
	return c
}

// Send an email to a receiver, with subject and body
func (c Client) Send(m Message) error {

	// validate message input
	if err := validateMessage(m); err != nil {
		return err
	}

	// construct message
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", c.SenderGmailAddress, m.Receiver, m.Subject, m.Body)
	auth := smtp.PlainAuth("", c.SenderGmailAddress, c.SenderPassword, mailServerHost)

	// send message
	if err := smtp.SendMail(c.mailServer, auth, c.SenderGmailAddress, []string{m.Receiver}, []byte(msg)); err != nil {
		return err
	}

	return nil
}

func validateMessage(m Message) error {
	if len(m.Body) < 1 {
		return errInvalidBody
	}

	if len(m.Subject) < 1 {
		return errInvalidSubject
	}

	if !emailValidator.MatchString(m.Receiver) {
		return errInvalidEmail
	}

	return nil
}
