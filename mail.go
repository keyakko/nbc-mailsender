/* 参考 https://qiita.com/Neetless/items/6c802dddc52fa09aaede */


package main

import (
    "log"
    "os"
    "fmt"

    "net/smtp"
)

func main() {
    m := mail{
        from:     "keyakko.dev@gmail.com",
        username: "keyakko.dev@gmail.com",
        password: "hogehoge",
        to:       "keyakko.dev@gmail.com",
        sub:      "クソメール",
        msg:      "hogehogehogehogehogehogehogehogehogehogehogehogehogehoge",
    }
    
    for ;; {
	    if err := gmailSend(m); err != nil {
            log.Println(err)
            os.Exit(1)
	    break
        }
	fmt.Println("send")
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
    smtpSvr := "127.0.0.1:1025"
    auth := smtp.PlainAuth("", m.username, m.password, "127.0.0.1")
    if err := smtp.SendMail(smtpSvr, auth, m.from, []string{m.to}, []byte(m.body())); err != nil {
        return err
    }
    return nil
}

