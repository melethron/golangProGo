package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Flow

	fmt.Println("Flow control")
	kayakPrice := 275.00

	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater, than 500, scopedVar is", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else if kayakPrice > 400 && kayakPrice < 500 {
		scopedVar := false
		fmt.Println("Price is between 400 and 500, scopedVar is", scopedVar)
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}

	// Using initialization statement
	priceString := "275"
	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println()
	// For loops
	fmt.Println("For loops")
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}
	fmt.Println()

	counter = 0
	// Same as before
	for counter <= 3 {
		fmt.Println("Counter:", counter)
		counter++
	}
	fmt.Println()

	// Initialization and completion statements with continuing loops
	fmt.Println("Initialization and completion statements")
	for i := 0; i <= 3; i++ {
		if i == 2 {
			fmt.Println("Continue")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// Enumerating sequences
	product := "Kayak"
	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}
	fmt.Println()

	// Only index
	for index := range product {
		fmt.Println("Index:", index)
	}
	fmt.Println()

	// Only character
	for _, char := range product {
		fmt.Println("Character: ", string(char))
	}
	fmt.Println()

	// Enumerating slice
	products := []string{"Kayak", "Lifejacker", "Soccer ball"}
	for i, p := range products {
		fmt.Println("Products index:", i, "Product:", p)
	}
	fmt.Println()

	// Switch statements
	fmt.Println("Using switch statements")
	product = "KayakA"
	for i, c := range product {
		switch c {
		case 'K', 'k':
			if c == 'k' {
				fmt.Println("Lowercase k at position", i)
				break
			}
			fmt.Println("K at position", i)
		case 'a':
			fmt.Println("A character")
			fallthrough
		case 'y':
			fmt.Println("y at position", i)
		default:
			fmt.Printf("String %c at position %d\n", c, i)
		}
	}

	// Init statement in switch
	for counter = 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non prime value:", val)
		}
	}
	fmt.Println()

	// Labels
	counter = 0
target:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto target
	}
	fmt.Println()
}
