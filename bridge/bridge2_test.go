package bridge

import (
	"testing"
)

func Test_bridge2(t *testing.T) {
	abstraction := RefineAbstraction{}
	abstraction.setImplementor(&ConcreateImplementorA{})
	abstraction.Operation()
}
