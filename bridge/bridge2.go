package bridge

import (
	"fmt"
)

//Implementor
type Implementor interface {
	Operator()
}

type ConcreateImplementorA struct{}

func (this *ConcreateImplementorA) Operator() {
	fmt.Println("我是ConcreateImplementorA")
}

type ConcreateImplementorB struct{}

func (this *ConcreateImplementorB) Operator() {
	fmt.Println("我是ConcreateImplementorB")
}

type Abstraction struct {
	implementor Implementor
}

func (this *Abstraction) setImplementor(implementor Implementor) {
	this.implementor = implementor
}
func (this *Abstraction) Operation() {
	fmt.Println("父类")
	this.implementor.Operator()
}

type RefineAbstraction struct {
	Abstraction
}

func (this *RefineAbstraction) Operation() {
	fmt.Println("子类")
	this.Abstraction.implementor.Operator()
}
