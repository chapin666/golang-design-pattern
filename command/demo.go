package command

import (
	"fmt"
)

// command
type Command interface {
	ExecuteCommand()
}
type Base struct {
	barbecuer Barbecuer
}

// cmd concrete
// mutton
type BakeMuttonCmd struct {
	Base
}

func NewBakeMutton(barbecuer Barbecuer) BakeMuttonCmd {
	return BakeMuttonCmd{Base{barbecuer}}
}
func (this *BakeMuttonCmd) ExecuteCommand() {
	this.barbecuer.BakeMutton()
}

//chicken
type BakeChickenCmd struct {
	Base
}

func NewChicken(barbecuer Barbecuer) BakeChickenCmd {
	return BakeChickenCmd{Base{barbecuer}}
}
func (this *BakeChickenCmd) ExecuteCommand() {
	this.barbecuer.BakeChickenWing()
}

// Receiver
type Barbecuer struct{}

func (this *Barbecuer) BakeMutton() {
	fmt.Println("烤羊肉")
}
func (this *Barbecuer) BakeChickenWing() {
	fmt.Println("烤鸡肉")
}

type Waiter struct {
	//cmd Command
	order map[Command]Command
}

func NewWaiter() Waiter {
	return Waiter{order: make(map[Command]Command)}
}

//下命令
// Executor
func (this *Waiter) addCommand(cmd Command) {
	//this.cmd = cmd
	this.order[cmd] = cmd
}

func (this *Waiter) ExecuteCmd() {
	for _, val := range this.order {
		val.ExecuteCommand()
	}
	//this.cmd.ExecuteCommand()
}
