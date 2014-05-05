package facade

import (
	"testing"
)

func TestFacede(t *testing.T) {
	fund := Fund{}
	fund.BuyStock()
	fund.SellStock()
}
