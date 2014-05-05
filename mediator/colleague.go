package mediator

import (
	"fmt"
)

type Colleague struct {
	mediator Mediator
}

// ConcreteColleagueA
type ConcreteColleagueA struct {
	Colleague
}

func NewConcreteColleagueA(media Mediator) ConcreteColleagueA {
	return ConcreteColleagueA{Colleague: Colleague{media}}
}

func (this *ConcreteColleagueA) sendMsg(msg string) {
	this.mediator.send(msg, "A")
	//fmt.Println(" ConcreteColleagueA send message ", msg)
}

func (this *ConcreteColleagueA) notify(msg string) {
	fmt.Println(" ConcreteColleagueA receive message ", msg)
}

// ConcreteColleagueB
type ConcreteColleagueB struct {
	Colleague
}

func NewConcreteColleagueB(media Mediator) ConcreteColleagueB {
	return ConcreteColleagueB{Colleague: Colleague{media}}
}

func (this *ConcreteColleagueB) sendMsg(msg string) {
	//fmt.Println(" ConcreteColleagueB send message ", msg)
	this.mediator.send(msg, "B")
}

func (this *ConcreteColleagueB) notify(msg string) {
	fmt.Println(" ConcreteColleagueB receive message ", msg)
}
