package singleton

import (
	"fmt"
	"testing"
)

func TestMutex(t *testing.T) {
	person1 := Newperson()
	person1.name = "程斌"

	person2 := Newperson()
	person2.name = "哈哈"

	fmt.Println(person1 == person2)
	fmt.Println(person1.name)
}
