package verification

import (
	"github.com/wneessen/go-mail"
	"log"
)

func SendMail(email string){
	m := mail.NewMsg()

	if err := m.From("myCompagnieMail@gmail.com"); err != nil{
		log.Fatalf(err.Error())
	}

	if err := m.To(email); err != nil{
		log.Fatalf(err.Error())
	}

	m.Subject("first mail for my website")
	m.SetBodyString(mail.TypeTextPlain,"hello world")
	c, err := mail.NewClient()
}