package notify

import (
	"bytes"
	"net/smtp"
)

func SendMail(to_addr string, from_addr string, host string, subject string, body string) error {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(host)
	if err != nil {
		return err
	}
	// Set the sender and recipient.
	c.Mail(from_addr)
	c.Rcpt(to_addr)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()
	buf := bytes.NewBufferString("Subject: " + subject + "\r\n" + "To: " + to_addr + "\r\n\r\n" + body)
	if _, err = buf.WriteTo(wc); err != nil {
		return err
	}
	return nil
}
