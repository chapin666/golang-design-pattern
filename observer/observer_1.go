package main

import (
	"container/list"
	"fmt"
)

type Secretary struct {
	action    string
	observers *list.List
}

func (this *Secretary) NewSecretary() {
	this.observers = list.New()
}

func (this *Secretary) Attach(stockObserver StockObserver) {
	this.observers.PushBack(stockObserver)
}

func (this *Secretary) Notify() {
	for e := this.observers.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(StockObserver); ok {
			v.update()
		}
	}
}
func (this *Secretary) secretaryAction() string {
	return this.action
}
func (this *Secretary) setSecretaryAction(action string) {
	this.action = action
}

//  observers
type StockObserver struct {
	name      string
	secretary *Secretary
}

func (this *StockObserver) NewStockObserver(name string, secretary *Secretary) {
	this.name = name
	this.secretary = secretary
}

func (this *StockObserver) update() {
	fmt.Println("name=", this.name, "secretary=", this.secretary.secretaryAction(), "..")
}

func main() {
	//前台
	zhuli := &Secretary{}
	zhuli.NewSecretary()

	//
	one := StockObserver{}
	one.NewStockObserver("程斌", zhuli)

	two := StockObserver{}
	two.NewStockObserver("张三2", zhuli)

	zhuli.Attach(one)
	zhuli.Attach(two)

	//Notice
	zhuli.setSecretaryAction("老板回来了")
	zhuli.Notify()
}
