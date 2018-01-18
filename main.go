package main

import (
	"bytes"
	"time"
	"log"
	"net/smtp"
)

func main() {
	for ;; {
		// Connect to the remote SMTP server.
		c, err := smtp.Dial("127.0.0.1:1025")
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()
		// Set the sender and recipient.
		c.Mail("keyakko.dev@gmail.com")
		c.Rcpt("keyakko.dev@gmail.com")
		// Send the email body.
		wc, err := c.Data()
		if err != nil {
			log.Fatal(err)
		}
		defer wc.Close()
		buf := bytes.NewBufferString("This is the email body.")
		if _, err = buf.WriteTo(wc); err != nil {
			log.Fatal(err)
		}
		c.Quit()
		time.Sleep(200000000)
	}
}
