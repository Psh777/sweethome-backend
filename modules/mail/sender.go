package mail

import (
	"../config"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

func getDialer() *gomail.Dialer {
	c := config.GetMyConfig().Env
	d := gomail.NewDialer(c.MailUrl, c.MailPort, c.MailUser, c.MailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}


func SendMail(a Mailer) {
	c := config.GetMyConfig().Env
	m := gomail.NewMessage()
	m.SetHeader("From", a.From)
	m.SetAddressHeader("From", a.From, c.MailNameFrom)
	m.SetHeader("To", a.To)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", a.Subject)
	m.SetBody("text/html", a.Body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := getDialer()
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("Error send mail:", err)
	} else {
		//fmt.Println("Send mail:", a.From, a.To, a.Subject, a.Body)
	}
}
