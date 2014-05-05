package main

import (
	"fmt"
	"strings"
)

type Company interface {
	add(c Company)
	remove(c Company)
	display(layer int)
	lineOfDuty()
}

type Base struct {
	name string
}
type CompanyA struct {
	Base
	list map[Company]Company
}

func NewCompanyA(name string) *CompanyA {
	list := make(map[Company]Company)
	return &CompanyA{Base{name}, list}
}
func (this *CompanyA) add(c Company) {
	this.list[c] = c
}
func (this *CompanyA) remove(c Company) {
	delete(this.list, c)
}
func (this *CompanyA) display(layer int) {
	fmt.Println(strings.Repeat("-", layer), "  ", this.name)

	for _, val := range this.list {
		val.display(layer + 2)
	}
}

func (this *CompanyA) lineOfDuty() {
	for _, val := range this.list {
		val.lineOfDuty()
	}
}

type CompanyB struct {
	Base
}

func NewCompanyB(name string) *CompanyB {
	return &CompanyB{Base{name}}
}
func (this *CompanyB) add(c Company) {

}
func (this *CompanyB) remove(c Company) {

}
func (this *CompanyB) display(layer int) {
	fmt.Println(strings.Repeat("-", layer), "  ", this.name)
}

func (this *CompanyB) lineOfDuty() {
	fmt.Println("CompanyB")
}

type CompanyC struct {
	Base
}

func NewCompanyC(name string) *CompanyC {
	return &CompanyC{Base{name}}
}
func (this *CompanyC) add(c Company) {

}
func (this *CompanyC) remove(c Company) {

}
func (this *CompanyC) display(layer int) {
	fmt.Println(strings.Repeat("-", layer), "  ", this.name)
}

func (this *CompanyC) lineOfDuty() {
	fmt.Println("CompanyC")
}

func main() {
	c := NewCompanyA("总部")
	c.add(NewCompanyB("分部A"))
	c.add(NewCompanyC("分步B"))

	c.display(2)
	c.lineOfDuty()
}
