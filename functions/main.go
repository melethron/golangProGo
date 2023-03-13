package main

import "fmt"

// Function with parameters
func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "tax:", taxAmount)
}

// omiting function parameters
func printPrice2(product string, price float64, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "tax:", taxAmount)
}

// Omiting all parameters
func printPrice3(string, float64, float64) {
	fmt.Println("No parameters")
}

// Variatic parameters
func printSuppliers(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func printSuppliers2(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

// No pointers
func swapValues(first, second int) {
	fmt.Println("Before swap:", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}

// Pointers as function parameters
func swapValues2(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

// Function returning results
func calcTax(price float64) float64 {
	return price + price*0.2
}

// Returning two results
func swapValues3(first, second int) (int, int) {
	return second, first
}

func calcTax2(price float64) (float64, bool) {
	if price > 100 {
		return price + price*0.2, true
	}
	return 0, false
}

// using named results
func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, taxDue := calcTax2(price); taxDue {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

// Defer keyword
func calcTotalPrice2(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call after 'function started'")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call before about to return")
	fmt.Println("function is about to return")
	return
}

func main() {
	fmt.Println("Functions")
	// Invoking function with parameters
	printPrice("kayak", 275.00, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)

	// Invoking function with omitted parameter
	printPrice2("kayak", 275.00, 0.2)
	printPrice2("Lifejacket", 48.95, 0.2)

	// Invoking function with all parameters omitted
	printPrice3("kayak", 275.00, 0.2)
	printPrice3("Lifejacket", 48.95, 0.2)
	fmt.Println()

	fmt.Println("Variatic parameters")
	// Invoking variatic parameters function
	printSuppliers("Kayak", "Acme kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers("Lifejacket", "Sail Safe Co")
	// Invoking variatic parameter function with no arguments for variatic parameter
	printSuppliers("Soccer ball")
	printSuppliers2("Soccer ball")
	printSuppliers2("Kayak", "Acme kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers2("Lifejacket", "Sail Safe Co")

	// Slices as values for periodic parameters
	names := []string{"Acme kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers2("Acme", names...)
	fmt.Println()

	fmt.Println("Using pointers as function parameters")
	val1, val2 := 10, 20
	fmt.Println("before calling the function", val1, val2)
	swapValues(val1, val2)
	fmt.Println("After calling the function", val1, val2)
	fmt.Println("Same funcion, using pointers")
	fmt.Println("before calling the function", val1, val2)
	swapValues2(&val1, &val2)
	fmt.Println("After calling the function", val1, val2)
	fmt.Println()

	fmt.Println("Defining and using function results")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		fmt.Println("Product:", product, "Price:", price, "Price with tax:", calcTax(price))
	}
	fmt.Println()

	// returning multiple results
	val1, val2 = 10, 20
	fmt.Println("Before swap:", val1, val2)
	val1, val2 = swapValues3(val1, val2)
	fmt.Println("After swap:", val1, val2)

	// using two values
	products2 := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products2 {
		if taxAmount, taxDue := calcTax2(price); taxDue {
			fmt.Println("Product:", product, "Price:", price, "Price with tax:", taxAmount)
		} else {
			fmt.Println("Product:", product, "Price:", price, "No tax due")
		}
	}

	// named results
	products3 := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	total1, tax1 := calcTotalPrice(products3, 10)
	fmt.Println("Total1:", total1, "Tax1:", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total2:", total2, "Tax2:", tax2)
	fmt.Println()

	// defer
	products4 := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	_, total := calcTotalPrice2(products4)
	fmt.Println("total:", total)
}
