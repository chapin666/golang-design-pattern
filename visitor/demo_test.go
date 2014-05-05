package visitor

import (
	"testing"
)

func TestDemo(t *testing.T) {
	objstruct := NewObjStructure()

	// add man and woman.
	objstruct.Attach(&Man{})
	objstruct.Attach(&Woman{})

	// show state
	action := Success{}
	objstruct.Display(&action)

	action2 := Failing{}
	objstruct.Display(&action2)
}
