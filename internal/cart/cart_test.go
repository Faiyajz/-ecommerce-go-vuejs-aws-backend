package cart

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestCart_TotalPriceVATInc(test *testing.T) {
	//GIVEN

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

	//WHEN
	actualToltalPrice, err := cart.TotalPriceVATInc()

	//THEN
	assert.NoError(test, err, "impossible to compute total price VAT included")

	expectedTotalPrice := money.New(100, "EUR")

	assert.Equal(test, expectedTotalPrice, actualToltalPrice)
}
