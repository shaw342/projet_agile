package verification

import (
	"github.com/wneessen/go-mail"
	"log"
)

func SendMail(email string){
	m := mail.NewMsg()

	if err := m.From("baruashawan41@gmail.com"); err != nil{
		log.Fatalf(err.Error())
	}

	if err := m.To(email); err != nil{
		log.Fatalf(err.Error())
	}

	m.Subject("first mail for my website")
	m.SetBodyString(mail.TypeTextPlain,"hello world")
	c, err := mail.NewClient("smtp.example.com", mail.WithPort(25), mail.WithSMTPAuth(mail.SMTPAuthPlain),
	mail.WithUsername("my_username"), mail.WithPassword("extremely_secret_pass"))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}

}


