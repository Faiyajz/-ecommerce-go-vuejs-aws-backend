package main

import "fmt"

type Product struct {
	ID        string
	Reference string
	Title     string
	Category
}

type Category struct {
	ID   string
	Name string
}

func main() {

	cat := Category{
		ID:   "catId",
		Name: "My Category",
	}

	products := []Product{
		{
			ID:        "42",
			Reference: "myRef",
			Title:     "My First Product",
			Category:  cat,
		},
		{
			ID:        "43",
			Reference: "myRef",
			Title:     "My Second Product",
			Category:  cat,
		},
		{
			ID:        "44",
			Reference: "myRef",
			Title:     "My Third Product",
			Category:  cat,
		},
	}

	for _, product := range products {

		fmt.Printf("Product of id %s belongs to category of id %s (name of the category: %s)\n", product.ID, product.Category.ID, product.Category.Name)

	}

}
