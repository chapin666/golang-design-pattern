package prototype

import (
	"fmt"
)

type Photo struct {
	url string
}

type Resume struct {
	name     string
	sex      string
	age      int
	timeArea string
	company  string
	Photo
}

func (this *Resume) setName(name string) {
	this.name = name
}

func (this *Resume) setInfo(sex string, age int) {
	this.age = age
	this.sex = sex
}

func (this *Resume) setPhoto(url string) {
	this.url = url
}

func (this *Resume) setExpericence(timeArea string, company string) {
	this.timeArea = timeArea
	this.company = company
}

func (this *Resume) display() {
	fmt.Println("name=", this.name, "sex=", this.sex,
		"age=", this.age, "expericence=", this.timeArea,
		"company=", this.company, "photo=", this.Photo.url)
}

func (this *Resume) clone() *Resume {
	newObj := *this
	return &newObj
}
