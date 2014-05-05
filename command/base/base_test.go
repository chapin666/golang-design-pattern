package command

import (
	"testing"
)

func TestBase(t *testing.T) {
	receicer := Receiver{}

	cmd := NewConcreteCmd(receicer)

	invoker := Invoker{}
	invoker.setCommand(&cmd)
	invoker.ExecuteCommand()

}
