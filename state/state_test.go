package state

import (
	"testing"
)

func TestState(t *testing.T) {

	work := new(Work)
	work.NewWork()
	work.SetHour(13)
	work.WriteProgram()

}
