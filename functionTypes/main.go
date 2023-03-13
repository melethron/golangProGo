package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + price*0.2
}

func calcWithoutTax(price float64) float64 {
	return price
}

// Functions as arguments
func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

// Using functions as results
func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

// Function literal
func selectCalculator2(price float64) calcFunc {
	if price > 100 {
		return func(price float64) float64 {
			return price + price*0.2
		}
	}
	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}

// Function type alias
type calcFunc func(float64) float64

var prizeGiveaway = false

// Using function closure
func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if prizeGiveaway {
			return 0
		}
		if price > threshold {
			return price + price*rate
		}
		return price
	}
}

func main() {
	fmt.Println("Function types")

	products := map[string]float64{
		"Kayak":      275,
		"lifejacket": 48.95,
	}

	for product, price := range products {
		var calcFunc calcFunc
		fmt.Println("Function assigned:", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}

	// function as arguments
	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}

	// Function as result
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}

	// Function closure
	watersportsProducts := map[string]float64{
		"Kayak":      275.00,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	waterCalc := priceCalcFactory(100, 0.2)
	soccerCalc := priceCalcFactory(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}
}
