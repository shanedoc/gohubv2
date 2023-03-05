package sms

//定制sms包driver interface
type Driver interface {
	// 发送短信
	Send(phone string, message Message, config map[string]string) bool
}
