package absfactory

import (
	"testing"
)

func TestSenderMail(t *testing.T) {
	var factory Provider = &sendMailFactory{}
	sender := factory.produce()
	sender.send()
}

func TestSmsSender(t *testing.T) {
	factory := new(sendSmsFactory)
	sender := factory.produce()
	sender.send()
}
