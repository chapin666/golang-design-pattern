package base

import (
	"fmt"
)

type Request struct {
	reqType    string
	reqContent string
	reqNumber  int
}

func (this *Request) ReqType() string {
	return this.reqType
}
func (this *Request) SetReqType(reqType string) {
	this.reqType = reqType
}
func (this *Request) ReqContent() string {
	return this.reqContent
}
func (this *Request) setReqContent(reqContent string) {
	this.reqContent = reqContent
}

func (this *Request) ReqNumber() int {
	return this.reqNumber
}
func (this *Request) SetReqNumber(number int) {
	this.reqNumber = number
}

type Handler interface {
	HandleRequest(request int)
}

type BaseHandle struct {
	handle Handler
}

func (this *BaseHandle) SetHandle(handle Handler) {
	this.handle = handle
}

type ConcreteHandleA struct {
	BaseHandle
}

func (this *ConcreteHandleA) HandleRequest(request int) {
	if request > 0 && request <= 10 {
		fmt.Println("ConcreteHandleA")
	} else {
		//this.HandleRequest(request)
		this.handle.HandleRequest(request)
	}
}

type ConcreteHandleB struct {
	BaseHandle
}

func (this *ConcreteHandleB) HandleRequest(request int) {
	if request > 10 && request <= 20 {
		fmt.Println("ConcreteHandleB")
	} else {
		//this.HandleRequest(request)
		fmt.Println("我是父亲")
	}
}
