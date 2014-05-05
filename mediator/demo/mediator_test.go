package demo

import (
	"testing"
)

func Test_Colleague(t *testing.T) {

	//中介者
	mediator := UnitedNationsA{}

	// Colleague
	usa := NewUSA(mediator)
	iraq := NewIraq(mediator)

	mediator.setUSA(usa)
	mediator.setIraq(iraq)

	iraq.Declare("哈哈")
	usa.Declare("hengheng")
}
