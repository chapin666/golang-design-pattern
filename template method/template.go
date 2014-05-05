package main

import (
	"fmt"
	"reflect"
)

type Question struct {
	parent interface{}
}

func (this *Question) question1() {
	fmt.Println("1+1=", this.Answer1())
}
func (this *Question) question2() {
	fmt.Print("2*3=", this.Answer2())
}
func (this *Question) question3() {
	fmt.Println("3/2=", this.Answer3())
}

func (this *Question) Answer1() float64 {
	if this.parent == nil {
		panic("请先答题1")
	}
	value := reflect.ValueOf(this.parent)
	res := value.MethodByName("Answer1").Call(nil)
	return res[0].Float()
}
func (this *Question) Answer2() float64 {
	if this.parent == nil {
		panic("请先答题2")
	}
	value := reflect.ValueOf(this.parent)
	res := value.MethodByName("Answer2").Call(nil)
	return res[0].Float()
}
func (this *Question) Answer3() float64 {
	if this.parent == nil {
		panic("请先答题3")
	}
	value := reflect.ValueOf(this.parent)
	res := value.MethodByName("Answer3").Call(nil)
	return res[0].Float()
}

type TestPaper struct {
	Question
}

func (this *TestPaper) Answer1() float64 {
	return 2
}
func (this *TestPaper) Answer2() float64 {
	return 6
}
func (this *TestPaper) Answer3() float64 {
	return 1
}

func NewTestPaper() *Question {
	paper := new(Question)
	paper.parent = new(TestPaper)
	return paper
}

func main() {
	paper := NewTestPaper()
	paper.question1()
	paper.question2()
}
