package bridge

import (
	"fmt"
)

type HandsetSoft interface {
	run()
}

type HandGame struct{}

func (this *HandGame) run() {
	fmt.Println("运行手机游戏")
}

type AddressList struct{}

func (this *AddressList) run() {
	fmt.Println("运行手机通讯录")
}

// Phone
type BaseBrand interface {
	run()
}
type HandsetBrand struct {
	hand HandsetSoft
}

func (this *HandsetBrand) SetSoft(hand HandsetSoft) {
	this.hand = hand
}

type HandsetBrandA struct {
	HandsetBrand
}

func (this *HandsetBrandA) run() {
	this.HandsetBrand.hand.run()
}

type HandsetBrandB struct {
	HandsetBrand
}

func (this *HandsetBrandB) run() {
	//this.HandsetBrand.hand.run()
	this.hand.run()
}
