package main

import "fmt"

func main() {

	type Cart struct {
		ID   string
		Paid bool
	}

	cart := Cart{
		ID:   "1234",
		Paid: true,
	}

	cartPointer := &cart

	// fmt.Println(&cart)
	fmt.Println(&cartPointer)

	cartDeref := *cartPointer
	fmt.Println(&cartDeref)

	// fmt.Println(cartPointer)

	// fmt.Println(cartDeref)
}
