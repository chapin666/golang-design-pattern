package main

import (
	"container/list"
	"fmt"
)

var obserList *list.List = list.New()

type Subject interface {
	attach(observer ObServer)
	notify()
	action1()
}

type Boss struct {
	action string
}

func (this *Boss) action1() {
	this.action = "好好工作"
}

func (this *Boss) attach(observers ObServer) {
	obserList.PushBack(observers)
}

func (this *Boss) notify() {
	for e := obserList.Front(); e != nil; e = e.Next() {
		if value, ok := e.Value.(ObServer); ok {
			value.update()
		}
	}
}

type ObServer interface {
	update()
}

type StockObserver struct {
	name    string
	subject Subject
}

func (this *StockObserver) NewStockObserver(name string, subject Subject) {
	this.name = name
	this.subject = subject
}

func (this *StockObserver) update() {
	fmt.Println(this.name, "通知", this.subject.action1())
}

func main() {
	subject := Boss{}

	one := new(StockObserver)
	one.NewStockObserver("张三", subject)
	two := new(StockObserver)
	two.NewStockObserver("李四", subject)

	subject.attach(one)
	subject.attach(two)

	subject.notify()
}
