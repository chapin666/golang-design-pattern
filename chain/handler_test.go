package base

import (
	"testing"
)

func TestHandle(t *testing.T) {
	a := ConcreteHandleA{}
	b := ConcreteHandleB{}

	a.SetHandle(&b)

	a.HandleRequest(12)
	a.HandleRequest(2)
	a.HandleRequest(4)
	a.HandleRequest(23)
	a.HandleRequest(14)
}
