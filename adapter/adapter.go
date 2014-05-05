package adapter

import (
	"fmt"
)

type BasePlayer struct {
	name string
}
type Player interface {
	attack()
	defense()
}

// Fowards
type Fowards struct {
	BasePlayer
}

func NewFowards(name string) *Fowards {
	return &Fowards{BasePlayer{name}}
}
func (this *Fowards) attack() {
	fmt.Println("前锋:", this.name, "进攻")
}
func (this *Fowards) defense() {
	fmt.Println("前锋:", this.name, "防守")
}

//ForeignCenter
type ForeignCenter struct {
	name string
}

func (this *ForeignCenter) attack() {
	fmt.Println("中锋:", this.name, "进攻")
}
func (this *ForeignCenter) defense() {
	fmt.Println("中锋:", this.name, "防守")
}

// Guards
type Guards struct {
	BasePlayer
}

func NewGuads(name string) *Guards {
	return &Guards{BasePlayer{name}}
}
func (this *Guards) attack() {
	fmt.Println("后卫:", this.name, "进攻")
}
func (this *Guards) defense() {
	fmt.Println("后卫:", this.name, "防守")
}

// transfer
type Transfer struct {
	ForeignCenter
}

func NewTransfer(name string) *Transfer {
	return &Transfer{ForeignCenter{name}}
}
