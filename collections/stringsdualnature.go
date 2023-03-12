package main

import (
	"fmt"
	"strconv"
)

func stringsDualNature() {
	fmt.Println("Strings dual nature")

	var price []rune = []rune("€58.95")

	var currency string = string(price[0])
	var amountString string = string(price[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency:", string(currency))
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parase error:", parseErr)
	}
	fmt.Println()

	// Enumerating strings
	fmt.Println("Enumerating strings")
	var price2 string = "€58.95"
	for index, char := range price2 {
		fmt.Println("Index:", index, "char:", string(char))
	}
	fmt.Println()

	// Enumerating bytes in the strings
	fmt.Println("Enumerating bytes in the strings")
	var price3 string = "€58.95"
	for i, b := range []byte(price3) {
		fmt.Println("Index:", i, "value:", b)
	}
}
