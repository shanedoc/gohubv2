package sms

import (
	"gohubv2/pkg/config"
	"sync"
)

type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

// SMS 是我们发送短信的操作类
type SMS struct {
	Driver Driver
}

var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

func NewSMS() *SMS {
	internalSMS = &SMS{
		Driver: &Aliyun{},
	}
	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
