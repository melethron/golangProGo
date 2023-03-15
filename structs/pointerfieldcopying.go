package main

import "fmt"

func copyProduct(product *Product2) Product2 {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

func pointerFieldCopying() {
	fmt.Println("Pointer fields copying")
	acme := &Supplier{"Acme Co", "New York"}

	p1 := newProduct2("Kayak", "Watersports", 275, acme)
	p2 := *p1
	p3 := copyProduct(p1)

	p1.name = "Original kayak"
	p1.Supplier.name = "boatco"

	for _, p := range []Product2{*p1, p2, p3} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}
}
