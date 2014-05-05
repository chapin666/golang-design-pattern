package memento

import (
	"fmt"
)

// Originator
type GameRole struct {
	vit int //体力
	atk int //战斗力
	def int //防御力
}

func (this *GameRole) StateDisplay() {
	fmt.Println("角色当前状态：")
	fmt.Println("体力：", this.vit)
	fmt.Println("战斗力：", this.atk)
	fmt.Println("防御力:", this.def)
}
func (this *GameRole) InitState() {
	this.atk = 100
	this.vit = 100
	this.def = 100
}
func (this *GameRole) Fighting() {
	this.vit = 0
	this.atk = 0
	this.def = 0
}

func (this *GameRole) SaveState() GameMemento {
	return GameMemento{this.vit, this.atk, this.def}
}
func (this *GameRole) RecoveryState(gameMemento GameMemento) {
	this.vit = gameMemento.vit
	this.atk = gameMemento.atk
	this.def = gameMemento.def
}

// memento
type GameMemento struct {
	vit int //体力
	atk int //战斗力
	def int //防御力s
}

// caretaker
type GameCaretaker struct {
	gameMemento GameMemento
}
