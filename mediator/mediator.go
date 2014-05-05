package mediator

// Mediator
type Mediator interface {
	send(msg string, kinds string)
}

// ConcreteMediator
type ConcreteMediator struct {
	creteA ConcreteColleagueA
	creteB ConcreteColleagueB
}

func (this *ConcreteMediator) setCreteA(creteA ConcreteColleagueA) {
	this.creteA = creteA
}
func (this *ConcreteMediator) setCreteB(creteB ConcreteColleagueB) {
	this.creteB = creteB
}
func (this *ConcreteMediator) send(msg string, kinds string) {
	if kinds == "A" {
		this.creteB.notify(msg)
	} else {
		this.creteA.notify(msg)
	}
}
