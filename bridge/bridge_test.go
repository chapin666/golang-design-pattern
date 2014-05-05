package bridge

import (
	"testing"
)

func Test_Bridg(t *testing.T) {

	var game HandsetSoft = &HandGame{}
	//list := AddressList{}
	//game := HandGame{}

	baseBrand := new(HandsetBrandA)
	baseBrand.SetSoft(game)
	baseBrand.run()

	brandB := new(HandsetBrandB)
	brandB.SetSoft(game)
	brandB.run()
}
