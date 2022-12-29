package cart

import (
	"testing"

	"github.com/Rhymond/go-money"
)

func TestCart_TotalPriceVATInc(test *testing.T) {

	// GIVEN
	items := []Item{
		{
			ID:               "42",
			ShortDescription: "A pair of socks",
			UnitPriceVATInc:  money.New(100, "EUR"),
			UnitPriceVATExc:  money.New(50, "EUR"),
			VAT:              money.New(50, "EUR"),
			Quantity:         1,
		},
	}
	cart := Cart{
		ID:           "42",
		CurrencyCode: "EUR",
		Items:        items,
	}
	// WHEN
	actualTotalPrice, err := cart.TotalPriceVATInc()

	// THEN
	if err != nil {
		test.Fail()
	}

	expectedTotalPrice := money.New(100, "EUR")
	isEqual, err := expectedTotalPrice.Equals(actualTotalPrice)
	if err != nil {
		test.Fail()
	}
	if !isEqual {
		test.Fail()
	}

}
