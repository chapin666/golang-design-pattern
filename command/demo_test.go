package command

import (
	"testing"
)

func TestDemo(t *testing.T) {
	//烧烤者
	barbecuer := Barbecuer{}

	//烤羊肉
	mut := NewBakeMutton(barbecuer)
	// 烤鸡肉
	chk := NewChicken(barbecuer)

	waiter := NewWaiter()
	//waiter.SetCommand(&mut)
	//waiter.ExecuteCmd()

	//waiter.SetCommand(&chk)
	//waiter.ExecuteCmd()

	waiter.addCommand(&mut)
	waiter.addCommand(&chk)
	waiter.ExecuteCmd()
}
