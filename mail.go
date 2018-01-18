package main

import (
    "log"
    "os"

    "sandspace/smtp_test/smtp"
)

func main() {
    m := mail{
        from:     "XXXXX@gmail.com",
        username: "XXXXX@gmail.com",
        password: "XXXXX",
        to:       "XXXXX@hotmail.co.jp",
        sub:      "test sub",
        msg:      "test msg",
    }

    if err := gmailSend(m); err != nil {
        log.Println(err)
        os.Exit(1)
    }
}

type mail struct {
    from     string
    username string
    password string
    to       string
    sub      string
    msg      string
}

func (m mail) body() string {
    return "To: " + m.to + "\r\n" +
        "Subject: " + m.sub + "\r\n\r\n" +
        m.msg + "\r\n"
}

func gmailSend(m mail) error {
    smtpSvr := "smtp.gmail.com:587"
    auth := smtp.PlainAuth("", m.username, m.password, "smtp.gmail.com")
    if err := smtp.SendMail(smtpSvr, auth, m.from, []string{m.to}, []byte(m.body())); err != nil {
        return err
    }
    return nil
}

