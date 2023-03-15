package main

import "fmt"

func comparingStructValues() {
	fmt.Println("Comparing struct values")

	type Product struct {
		name, category string
		price          float64
	}
	p1 := Product{"Kayak", "Watersports", 279}
	p2 := Product{"Kayak", "Watersports", 279}
	p3 := Product{"Kayak", "Boats", 275.00}

	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	fmt.Println("p3", p3)
	fmt.Println("p1==p2", p1 == p2)
	fmt.Println("p2==p3", p2 == p3)
	fmt.Println()

	fmt.Println("Converting between structs")
	type Item struct {
		name, category string
		price          float64
	}
	pi := Item(p1)
	fmt.Println("pi:", pi)
	fmt.Println()

}
