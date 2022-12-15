package main

import (
	"fmt"
	"log"
)

type Product struct {
	ID          string
	Description string
	Stock       uint
}

func (p Product) IsInStock() bool {
	if p.Stock > 0 {
		return true
	}
	return false
}

func (p *Product) SetDescription(description string) error {
	if len(p.Description) <= 10 || len(p.Description) >= 250 {
		p.Description = description
		return nil
	} else {
		return fmt.Errorf("length of description is not correct: got length %d", len(description))
	}
}

func main() {
	product := Product{
		ID:    "868",
		Stock: 10,
	}

	fmt.Println(product.IsInStock())
	err := product.SetDescription("FaiyajFaiyajFaiyaj")

	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("Success")
}
