package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := list.New()
	a.PushBack(1)
	a.PushBack("test")
	a.PushBack("hah")
	a.PushBack(false)

	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Println(i)
	}
}
