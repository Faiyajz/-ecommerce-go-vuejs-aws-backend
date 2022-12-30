package cart

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestCart_TotalPriceVATInc(test *testing.T) {
	test.Run(
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
	)
}


func TestCart_TotalPriceVATInc2(test *testing.T) {
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
		{
			ID:               "43",
			ShortDescription: "A T-shirt with a small gopher",
			UnitPriceVATInc:  money.New(3480, "EUR"),
			UnitPriceVATExc:  money.New(2900, "EUR"),
			VAT:              money.New(580, "EUR"),
			Quantity:         2,
		},
	}

	cart := Cart{
		ID:           "42",
		CurrencyCode: "EUR",
		Items:        items,
	}

	//WHEN
	actualToltalPriceVATInc, err := cart.TotalPriceVATInc()

	//THEN
	assert.NoError(test, err, "should have no error when total VAT inc is computed")

	expectedTotalPriceVATINC := money.New(7060, "EUR")

	assert.Equal(test, expectedTotalPriceVATINC, actualToltalPriceVATInc)
}

func TestCart_TotalPriceVATInc3(test *testing.T) {
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
	_, err := cart.TotalPriceVATInc()

	//THEN
	assert.NoError(test, err, "When I add an item with a currency X to a basket Y the method TotalPriceVATINC should Fail")

}
