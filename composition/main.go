package main

import (
	"composition/store"
	"fmt"
)

func main() {
	fmt.Println("Hello, composition")

	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, true),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.23, 2, true),
	}

	for _, b := range boats {
		fmt.Println("Name:", b.Name, "Price:", b.Price(0.2))
	}

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}

	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
	}

	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend special", product, 50)

	Name, price, Price := deal.GetDetails()

	fmt.Println("name:", Name)
	fmt.Println("price field:", price)
	fmt.Println("Price method:", Price)

	products := map[string]store.ItemForSale{
		"kayak": store.NewBoat("kayak", 279, 1, false),
		"ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}

	for key, p := range products {
		fmt.Println("Key:", key, "Price:", p.Price(0.2))
	}

	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("name:", item.GetName(), "category:", item.GetCategory(), "Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

}
