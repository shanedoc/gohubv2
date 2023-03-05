package mail

import (
	"gohubv2/pkg/config"
	"sync"
)

//mail 发送对象

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Plaintext message (optional)
	HTML    []byte // Html message (optional)
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

//NewMailer单例 获取实例对象
func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})
	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
