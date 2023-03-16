package main

import (
	"fmt"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	fmt.Println("hello, packages!")

	product := *store.NewProduct("Kayak", "Watersports", 275)
	product.SetPrice(100)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{product},
	}

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))

	fmt.Println(cart.CustomerName)
	fmt.Println(cart.GetTotal())

}
