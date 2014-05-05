package absfactory

import (
	"fmt"
)

// base interface
type Sender interface {
	send()
}

// 以下两个结构体是Sender的具体实现
// 1.MailSneder
type MailSender struct{}

func (this *MailSender) send() {
	fmt.Println("发送电子邮件")
}

// 2.SmsSender
type SmsSender struct{}

func (this *SmsSender) send() {
	fmt.Println("发送手机信息")
}

// base for factory
type Provider interface {
	produce() Sender
}

// Provider的具体实现
type sendSmsFactory struct{}

func (this *sendSmsFactory) produce() Sender {
	return new(SmsSender)
}

type sendMailFactory struct{}

func (this *sendMailFactory) produce() Sender {
	return new(MailSender)
}
