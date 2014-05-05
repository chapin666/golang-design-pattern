package singleton

import (
	"sync"
)

var p *person

type person struct {
	name string
}

func Newperson() *person {
	m := new(sync.Mutex)

	if p == nil {
		m.Lock()
		if p == nil {
			p = new(person)
		}
		m.Unlock()
	}

	return p
}
