package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func calcTax(p *Product) {
	if p.price > 100 {
		p.price += p.price * 0.2
	}
}

func calcTax2(p *Product) *Product {
	if p.price > 100 {
		p.price += p.price * 0.2
	}
	return p
}

func Newproduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func structsAndPointers() {
	fmt.Println("Understanding structs and pointers")

	p1 := Product{"Kayak", "Watersports", 279}
	p2 := p1
	p1.name = "Original Kayak"
	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", p2.name)

	p3 := Product{"Kayak", "Watersports", 279}
	p4 := &p3
	p3.name = "Original Kayak"
	fmt.Println("P3:", p3.name)
	fmt.Println("P4:", (*p4).name)
	fmt.Println()

	// Struct pointer convenience syntax
	fmt.Println("Struct pointer convenience syntax")
	p5 := &Product{"Kayak", "Watersports", 279}
	calcTax(p5)
	fmt.Println("name:", p5.name, "category:", p5.category, "price:", p5.price)
	p6 := calcTax2(&Product{"Kayak", "Watersports", 279})
	fmt.Println("name:", p6.name, "category:", p6.category, "price:", p6.price)
	fmt.Println()

	fmt.Println("Struct constructor functions")
	products := [2]*Product{
		Newproduct("Kayak", "Watersports", 275),
		Newproduct("Hat", "Skiing", 42.50),
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
	}
	fmt.Println()

}
