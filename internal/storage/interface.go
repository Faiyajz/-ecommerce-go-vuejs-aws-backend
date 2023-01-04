package storage

import "backend/internal/product"

type Storage interface {
	CreateProduct(product product.Product) error
	// Foo()
}
