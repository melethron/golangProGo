package main

import (
	"fmt"
	"sort"
)

func maps() {
	fmt.Println("Maps")
	fmt.Println("Creating a map")
	productsMap := make(map[string]float64, 10)
	productsMap["Kayak"] = 279
	productsMap["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(productsMap))
	fmt.Println("Price:", productsMap["Kayak"])
	fmt.Println("Price:", productsMap["Lifejacket"])
	fmt.Println()

	// Map literal Syntax
	fmt.Println("Map literal syntax")
	productsMap2 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	fmt.Println("Map size:", len(productsMap2))
	fmt.Println("Price:", productsMap2["Kayak"])
	fmt.Println("Price:", productsMap2["Lifejacket"])
	fmt.Println()

	// Checking for items in map
	fmt.Println("Checking for item in map")
	productsMap3 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	fmt.Println("Before adding hat", productsMap3["Hat"])
	productsMap3["Hat"] = 0
	fmt.Println("After adding hat", productsMap3["Hat"])
	val, ok := productsMap3["Hat"]
	if ok {
		fmt.Println("Value is", val)
	} else {
		fmt.Println("No stored value")
	}
	fmt.Println()

	// Removing items from map
	fmt.Println("Remove item from map")
	productsMap4 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	delete(productsMap4, "Hat")
	if val, ok = productsMap4["Hat"]; ok {
		fmt.Println("Value is", val)
	} else {
		fmt.Println("No stored value")
	}
	fmt.Println()

	// Enumerating maps
	fmt.Println("Enumerating maps")
	productsMap5 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}

	for key, value := range productsMap5 {
		fmt.Println("Key:", key, "value:", value)
	}
	fmt.Println()

	// Enumerating map in order
	fmt.Println("Enumerating in order")
	productsMap6 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	keys := make([]string, 0, len(productsMap6))
	for key := range productsMap6 {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", productsMap6[key])
	}
	fmt.Println()
}
