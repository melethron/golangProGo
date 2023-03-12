package main

import "fmt"

func arrays() {
	// Arrays
	fmt.Println("Arrays")
	var names [3]string

	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"

	fmt.Println(names)

	names2 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names2)
	names3 := [...]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names3)

	otherArray := names
	names[0] = "Canoe"
	fmt.Println("Names:", names)
	fmt.Println("OtherArray:", otherArray)

	var otherArray2 = &names
	names[0] = "Kayak"
	fmt.Println("Names:", names)
	fmt.Println("otherArray2:", *otherArray2)

	// Comparing arrays
	fmt.Println("Comparing arrays")
	names = [3]string{"Kayak", "Lifejacket", "Paddle"}
	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	same := names == moreNames
	fmt.Println("Comparison:", same)

	// Enumerating array
	for i, name := range names {
		fmt.Printf("Index: %d, name: %s\n", i, name)
	}
	fmt.Println()
}
