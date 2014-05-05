package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {

	mediator := new(ConcreteMediator)

	creteA := NewConcreteColleagueA(mediator)
	creteB := NewConcreteColleagueB(mediator)

	mediator.setCreteA(creteA)
	mediator.setCreteB(creteB)

	creteA.sendMsg("吃饭没有")
	creteB.sendMsg("没有就好")
}
