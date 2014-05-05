package builder

import (
	"fmt"
)

// builder
type Person interface {
	buildHair()
	buildFace()
	buildBody()
	buildArm()
	buildLeg()
}

//ConcreteBuilder
type Boy struct{}

func (this *Boy) buildHair() {
	fmt.Println("短头发")
}
func (this *Boy) buildFace() {
	fmt.Println("男生脸")
}
func (this *Boy) buildBody() {
	fmt.Println("男生身体")
}
func (this *Boy) buildArm() {
	fmt.Println("男生手臂")
}
func (this *Boy) buildLeg() {
	fmt.Println("男生腿")
}

type Girl struct{}

func (this *Girl) buildHair() {
	fmt.Println("长头发")
}
func (this *Girl) buildFace() {
	fmt.Println("女生脸")
}
func (this *Girl) buildBody() {
	fmt.Println("女生身体")
}
func (this *Girl) buildArm() {
	fmt.Println("女生手臂")
}
func (this *Girl) buildLeg() {
	fmt.Println("女生腿")
}

//Director
type PersonDirector struct {
	p Person
}

func (this *PersonDirector) CreatePerson() {
	if this.p == nil {
		panic("请指定要构造的人")
	}
	this.p.buildHair()
	this.p.buildFace()
	this.p.buildBody()
	this.p.buildArm()
	this.p.buildLeg()
}
