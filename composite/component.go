package composite

import (
	"fmt"
	"strings"
)

type Component interface {
	add(Component)
	remove(Component)
	display(int)
}

type BaseComponent struct {
	name string
}
type ConcreateCompany struct {
	BaseComponent
	list map[Component]Component
}

func NewConcreateCompany(name string) *ConcreateCompany {
	list := make(map[Component]Component)
	return &ConcreateCompany{BaseComponent{name}, list}
}

func (this *ConcreateCompany) add(component Component) {
	this.list[component] = component
}
func (this *ConcreateCompany) remove(component Component) {
	delete(this.list, component)
}

func (this *ConcreateCompany) display(layer int) {
	fmt.Println(strings.Repeat("-", layer), " ", this.name)

	for _, val := range this.list {
		val.display(layer + 2)
	}
}

//Leaf
type Leaf struct {
	BaseComponent
}

func NewLeaf(name string) *Leaf {
	return &Leaf{BaseComponent{name}}
}
func (this *Leaf) add(component Component) {
	fmt.Println("我是叶子节点，不能添加")
}
func (this *Leaf) remove(component Component) {
	fmt.Println("我是叶子节点，不能删除")
}
func (this *Leaf) display(layer int) {
	fmt.Println(strings.Repeat("-", layer), this.name)
}
