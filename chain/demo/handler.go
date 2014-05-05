package demo

import (
	"fmt"
)

type Manager interface {
	RequestApplication(request Request)
}

type Base struct {
	name   string
	manage Manager
}

func (this *Base) SetManager(manage Manager) {
	this.manage = manage
}

type CommManager struct {
	Base
}

func NewCommManager(name string) CommManager {
	return CommManager{Base: Base{name: name}}
}

func (this *CommManager) RequestApplication(request Request) {
	if request.ReqType() == "请假" && request.ReqNumber() <= 3 {
		fmt.Println("我是经理，允许你请假小于3天")
	} else {
		this.manage.RequestApplication(request)
	}
}

type MajorManager struct {
	Base
}

func NewMajorManager(name string) MajorManager {
	return MajorManager{Base: Base{name: name}}
}

func (this *MajorManager) RequestApplication(request Request) {
	if request.ReqType() == "请假" && request.ReqNumber() <= 10 {
		fmt.Println("我是总监，允许你请假小于10天")
	} else {
		//this.manage.RequestApplication(request)
		fmt.Println("去找总经理吧")
	}
}
