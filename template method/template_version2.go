package main

import (
	"fmt"
)

type Question struct {
	Vir interface{}
}

func (this *Question) question1() {
	fmt.Println("1+1=", this.Vir.(IAnswer).answer1())
}
func (this *Question) question2() {
	fmt.Println("4*5=", this.Vir.(IAnswer).answer2())
}
func (this *Question) question3() {
	fmt.Println("6/3=", this.Vir.(IAnswer).answer3())
}

type IAnswer interface {
	answer1() float64
	answer2() float64
	answer3() float64
}
type TestPaperA struct {
	Question
}

func (this *TestPaperA) answer1() float64 {
	return 2
}
func (this *TestPaperA) answer2() float64 {
	return 20
}
func (this *TestPaperA) answer3() float64 {
	return 2
}

type TestPaperB struct {
	Question
}

func (this *TestPaperB) answer1() float64 {
	return 2
}
func (this *TestPaperB) answer2() float64 {
	return 20
}
func (this *TestPaperB) answer3() float64 {
	return 12
}

func NewTestPaperA() *Question {
	paper := new(Question)
	testA := new(TestPaperA)
	paper.Vir = testA

	return paper
}

func NewTestPaperB() *Question {
	paper := new(Question)
	paper.Vir = new(TestPaperB)

	return paper
}

func main() {
	paper := NewTestPaperA()
	paper.question1()

	paper = NewTestPaperB()
	paper.question2()
	paper.question3()
}
