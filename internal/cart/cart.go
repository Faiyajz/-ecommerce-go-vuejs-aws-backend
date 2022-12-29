package cart

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

type Cart struct {
	ID           string
	CurrencyCode string
	Items        []Item
}

type Item struct {
	ID               string
	ShortDescription string
	Quantity         uint
	UnitPriceVATExc  *money.Money
	VAT              *money.Money
	UnitPriceVATInc  *money.Money
}

func (cart Cart) TotalPrice() (*money.Money, error) {
	totalPrice := money.New(0, cart.CurrencyCode)

	for _, item := range cart.Items {
		itemPrice := item.UnitPriceVATInc.Multiply(int64(item.Quantity))

		var err error
		totalPrice, err = totalPrice.Add(itemPrice)
		if err != nil {
			return nil, fmt.Errorf("impossible to add item price to total price: %w", err)
		}
	}

	return totalPrice, nil
}
