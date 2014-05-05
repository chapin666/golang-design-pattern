package demo

import (
	"fmt"
)

// Colleague
// base
type Country struct {
	unit UnitedNations
}

// USA
type USA struct {
	Country
}

func NewUSA(mediator UnitedNations) USA {
	return USA{Country{mediator}}
}

func (this *USA) Declare(msg string) {
	this.unit.Declare(msg, this)
}

func (this *USA) getMessage(msg string) {
	fmt.Println(msg)
}

// Iraq
type Iraq struct {
	Country
}

func NewIraq(mediator UnitedNations) Iraq {
	return Iraq{Country{mediator}}
}

func (this *Iraq) Declare(msg string) {
	this.unit.Declare(msg, this)
}

func (this *Iraq) getMessage(msg string) {
	fmt.Println(msg)
}
