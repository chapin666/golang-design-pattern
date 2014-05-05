package facade

import (
	"fmt"
)

//SubSystem
type Stock1 struct{}

func (this *Stock1) Buy() {
	fmt.Println("买下第一支股")
}
func (this *Stock1) Sell() {
	fmt.Println("卖出第一支股")
}

type Stock2 struct{}

func (this *Stock2) Buy() {
	fmt.Println("买下第二支股")
}
func (this *Stock2) Sell() {
	fmt.Println("卖出第二支股")
}

type Stock3 struct{}

func (this *Stock3) Buy() {
	fmt.Println("买下第三支股")
}
func (this *Stock3) Sell() {
	fmt.Println("卖出第三支股")
}

type NationalDept struct{}

func (this *NationalDept) Buy() {
	fmt.Println("买下国债")
}
func (this *NationalDept) Sell() {
	fmt.Println("卖出国债")
}

// Facade
type Fund struct {
	stock1       Stock1
	stock2       Stock2
	stock3       Stock3
	nationalDept NationalDept
}

func (this *Fund) BuyStock() {
	this.stock1.Buy()
	this.stock2.Buy()
	this.stock3.Buy()
	this.nationalDept.Buy()
}

func (this *Fund) SellStock() {
	this.stock1.Sell()
	this.stock2.Sell()
	this.nationalDept.Buy()
}
