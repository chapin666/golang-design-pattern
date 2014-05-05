package main

import (
	"fmt"
)

// LeiFeng interface
type LeiFeng interface {
	Sweep()
	Wash()
	BuyRice()
}

//Student implements
type Student struct {
	name string
}

func (this *Student) Sweep() {
	fmt.Println(this.name, "打扫卫生")
}

func (this *Student) Wash() {
	fmt.Println(this.name, "洗衣服")
}

func (this *Student) BuyRice() {
	fmt.Println(this.name, "买米")
}

//Volunteer implements
type Volunteer struct {
	name string
}

func (this *Volunteer) Sweep() {
	fmt.Println(this.name, "打扫卫生")
}
func (this *Volunteer) Wash() {
	fmt.Println(this.name, "洗衣服")
}

func (this *Volunteer) BuyRice() {
	fmt.Println(this.name, "买米")
}

// 工厂方法接口
type ILeiFengFactory interface {
	CreateLeiFeng(name string) LeiFeng
}

type StudentFactory struct{}

func (this *StudentFactory) CreateLeiFeng(name string) LeiFeng {
	return &Student{name: name}
}

type VolunteerFactory struct{}

func (this *VolunteerFactory) CreateLeiFeng(name string) LeiFeng {
	return &Volunteer{name: name}
}

func main() {
	studentFactory := new(VolunteerFactory)
	student := studentFactory.CreateLeiFeng("志愿者程斌")
	student.BuyRice()
	student.Sweep()
	student.Wash()

}
