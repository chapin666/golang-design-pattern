package decorator

import (
	"fmt"
)

// Component
type Phone interface {
	showMe()
}

// ConcreteComponent
// Anroid
type Android struct {
	Phone
	version string
}

func (this *Android) showMe() {
	fmt.Println("我是Android系统，当前版本是：", this.version)
}

// IOS
type IOS struct {
	Phone
	version string
}

func (this *IOS) showMe() {
	fmt.Println("我是IOS系统，当前版本是：", this.version)
}

// Decorator
// Decorator for users.
type Decorator struct {
	Phone
}

func (this *Decorator) Decorate(phone Phone) {
	this.Phone = phone
}
func (this *Decorator) showMe() {
	if this.Phone != nil {
		this.Phone.showMe()
	}
}

// ConcreteDecorator
// Add DIY
type DIY struct {
	Decorator
}

func (this *DIY) showMe() {
	this.Phone.showMe()
	fmt.Println("增加手机挂件")
}

// Add screen protection
type SrcProtection struct {
	Decorator
}

func (this *SrcProtection) showMe() {
	this.Phone.showMe()
	this.addSrcProtection()
}

func (this *SrcProtection) addSrcProtection() {
	fmt.Println("增加屏保")
}
