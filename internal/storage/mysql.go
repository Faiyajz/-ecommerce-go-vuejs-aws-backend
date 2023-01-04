package storage

import "backend/internal/product"

type MySQL struct {
}

func (mysql *MySQL) CreateProduct(product product.Product) error {

	panic("implement me")

}

// func (mysql MySQL) Foo() {

// }

func (mysql MySQL) String() string {

	return "MYSQL"

}
