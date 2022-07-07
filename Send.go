package email

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

type Sender struct {
	Name            string   // tim
	Mail            string   // tim@gmail.com
	MailSmtpAddress string   // smtp.qq.com
	MailAuthcode    string   // hauosighiaus
	Subject         string   // Notice
	Text            string   // Hello
	Attach          []string // info.xlsx
}

type Recipient struct {
	Name string // tom
	Mail string // tom@gmail.com
}

func (s *Sender) Send(r Recipient) error {
	d := gomail.NewDialer(s.MailSmtpAddress, 465, s.Mail, s.MailAuthcode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(s.Mail, s.Name))
	m.SetHeader("To", m.FormatAddress(r.Mail, r.Name))
	m.SetHeader("Subject", s.Subject)
	m.SetBody("text/plain", s.Text)
	for _, f := range s.Attach {
		m.Attach(f)
	}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
