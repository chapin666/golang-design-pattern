package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	var a Player = NewFowards("张三")
	a.attack()
	var b Player = NewGuads("王五")
	b.attack()
	var c Player = NewTransfer("李四")
	c.attack()
	c.defense()
}
