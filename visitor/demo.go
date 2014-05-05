package visitor

import (
	"fmt"
)

// person interface
type person interface {
	accept(action Action)
}

// man and woman
type Man struct{}

func (this *Man) accept(action Action) {
	action.getManstate()
}

type Woman struct{}

func (this *Woman) accept(action Action) {
	action.getWomanstate()
}

// action interface (state)
type Action interface {
	getManstate()
	getWomanstate()
}

// success state
type Success struct{}

func (this *Success) getManstate() {
	fmt.Println("男人成功时，背后多半有一个成功的女人")
}
func (this *Success) getWomanstate() {
	fmt.Println("女人成功时，背后多半有一个失败的男人")
}

// failing state
type Failing struct{}

func (this *Failing) getManstate() {
	fmt.Println("男人失败时，背后多半有一个失败的女人")
}
func (this *Failing) getWomanstate() {
	fmt.Println("女人失败时，背后多半有一个失败的男人")
}

type ObjectStructure struct {
	persons map[person]person
}

func NewObjStructure() ObjectStructure {
	return ObjectStructure{persons: make(map[person]person)}
}

func (this *ObjectStructure) Attach(p person) {
	this.persons[p] = p
}

func (this *ObjectStructure) Detach(p person) {
	//this.persons[p]
}

func (this *ObjectStructure) Display(action Action) {
	for _, val := range this.persons {
		val.accept(action)
	}
}
