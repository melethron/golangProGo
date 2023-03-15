package main

import "fmt"

func definingStruct() {
	fmt.Println("Defining struct")

	type Product struct {
		name, category string
		price          float64
	}

	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println(kayak.price)
	fmt.Println()

	fmt.Println("Partially assigning struct values")
	kayak = Product{
		name:     "Kayak",
		category: "Watersports",
	}

	fmt.Println("price is not set")
	fmt.Println(kayak.name, kayak.category, kayak.price)

	// Using field positions to created struct values
	kayak = Product{"Kayak", "Watersports", 279}
	fmt.Println(kayak)
	fmt.Println()

	// Defining embedded fields
	fmt.Println("Embedded fields")

	/*
		type Product struct {
			name, category string
			price          float64
		}
	*/

	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}

	stockItem := StockLevel{
		Product:   Product{"Kayak", "Watersports", 279},
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     2,
	}
	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Altername name:", stockItem.Alternate.name)
	fmt.Println("Count:", stockItem.count)
	fmt.Println()
}
