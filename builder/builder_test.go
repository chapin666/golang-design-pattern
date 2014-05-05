package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	p := PersonDirector{p: &Boy{}}
	p.CreatePerson()
}
