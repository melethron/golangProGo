package main

import "fmt"

type Product2 struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func newProduct2(name, category string, price float64, supplier *Supplier) *Product2 {
	return &Product2{name, category, price, supplier}
}

func pointerStructFields() {
	fmt.Println("Using Pointer Types for Struct Fields")
	acme := &Supplier{"Acme co", "new York"}

	products := [2]*Product2{
		newProduct2("Kayak", "Watersports", 275, acme),
		newProduct2("Hat", "Skiing", 42.50, acme),
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}
}
