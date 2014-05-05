package command

import (
	"fmt"
)

// Command
type Command interface {
	Execute()
}
type Base struct {
	receiver Receiver
}

//ConcreteCommand
type ConcreteCommand struct {
	Base
}

func NewConcreteCmd(receiver Receiver) ConcreteCommand {
	return ConcreteCommand{Base: Base{receiver}}
}

func (this *ConcreteCommand) Execute() {
	this.receiver.Action()
}

//Invoker
type Invoker struct {
	command Command
}

func (this *Invoker) setCommand(command Command) {
	this.command = command
}
func (this *Invoker) ExecuteCommand() {
	this.command.Execute()
}

// Receiver
type Receiver struct{}

func (this *Receiver) Action() {
	fmt.Println("执行请求")
}
