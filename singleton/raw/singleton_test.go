package pattern

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	fmt.Println("normal testing:")
	singleton := NewSingleton("go home")
	singleton.doSometing()
	singleton = NewSingleton("go school")
	singleton.doSometing()
}

func TestGoSingleton(t *testing.T) {
	fmt.Println("goroutine testing:")

	flag := make(chan bool)
	go newObject(flag, "go home")
	go newObject(flag, "do homework")
	<-flag
	<-flag
	fmt.Println("finish goroutine")
}

func newObject(c chan bool, something string) {
	singleton := NewSingleton(something)
	singleton.doSometing()
	c <- false
}
