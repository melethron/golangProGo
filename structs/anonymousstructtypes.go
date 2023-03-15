package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("name:", val.name)
}

func anonymousStructtypes() {
	fmt.Println("Anonymous struct types")
	type Product struct {
		name, category string
		price          float64
	}

	type Item struct {
		name     string
		category string
		price    float64
	}

	prod := Product{name: "Kayak", category: "Watersports", price: 275}
	item := Item{name: "Stadium", category: "Soccer", price: 75000}

	writeName(prod)
	writeName(item)
	fmt.Println()

	// Anonymous example
	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println(builder.String())
	fmt.Println()
}
