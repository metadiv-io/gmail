package gmail

import (
	"crypto/tls"

	gmail "gopkg.in/mail.v2"
)

const TYPE_PLAIN_TEXT = "text/plain"
const TYPE_HTML = "text/html"

type Auth struct {
	Host     string
	Port     int
	User     string
	Password string // you can get your token from your google account
}

func NewAuth(host string, port int, user string, password string) *Auth {
	return &Auth{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}
}

type Message struct {
	From       string
	To         []string
	Cc         []string
	Subject    string
	Body       string
	BodyType   string // text/plain or text/html
	Attachment []string
}

func SendEmail(auth *Auth, mail *Message) error {
	msg := gmail.NewMessage()
	msg.SetHeader("From", mail.From)
	msg.SetHeader("To", mail.To...)
	msg.SetHeader("Cc", mail.Cc...)
	msg.SetHeader("Subject", mail.Subject)
	msg.SetBody(mail.BodyType, mail.Body)
	for _, attachment := range mail.Attachment {
		msg.Attach(attachment)
	}

	d := gmail.NewDialer(auth.Host, auth.Port, auth.User, auth.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(msg)
}
