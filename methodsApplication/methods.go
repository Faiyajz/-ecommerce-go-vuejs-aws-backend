package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/Rhymond/go-money"
)

type Item struct {
	ID    string
	Name  string
	Price *money.Money
}

type Cart struct {
	ID       string
	IsLocked bool
	Items    []Item
}

func (cart Cart) TotalPrice() (*money.Money, error) {

	totalPrice := money.New(0, "EUR")
	for _, item := range cart.Items {
		var err error
		totalPrice, err = totalPrice.Add(item.Price)
		if err != nil {
			return nil, fmt.Errorf("impossible to compute the total price: %w", err)
		}
	}
	return totalPrice, nil
}

func (cart *Cart) Lock() error {

	if cart.IsLocked {
		return errors.New("the cart is already locked, cannot be locked")
	}
	cart.IsLocked = true
	return nil
}

func main() {
	items := []Item{
		{
			ID:    "1",
			Name:  "Book",
			Price: money.New(1000, "EUR"),
		},
		{
			ID:    "868",
			Name:  "Book 2",
			Price: money.New(1200, "EUR"),
		},
	}
	cart := Cart{
		ID:    "42",
		Items: items,
		IsLocked: true,
	}
	total, err := cart.TotalPrice()

	if err != nil {
		log.Fatalf("Error")
	}

	fmt.Println(total.Display())
	err = cart.Lock()
	if err != nil {
		log.Fatalf("Error while locking the cart: %s", err)
	}
	fmt.Println("Cart is locked")
}
