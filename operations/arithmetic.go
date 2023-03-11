package main

import (
	"fmt"
	"math"
)

func arithmetic() {
	// Arithmetic operations
	price, tax := 275.0, 27.40

	fmt.Println("basic arithmetic")
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)
	fmt.Println()

	// Arithmetic overflow
	fmt.Println("Arithmetic overflow")
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64
	fmt.Println(2 * intVal)
	fmt.Println(2 * floatVal)
	fmt.Println(math.IsInf((floatVal * 2), 0))
	fmt.Println()

	// Reminder operator
	fmt.Println("Reminder operator")
	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))
	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)
	fmt.Println()

	// Increment and decrement operators
	fmt.Println("Increment and decrement operators")
	value := 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value--
	fmt.Println(value)
	fmt.Println()
}
