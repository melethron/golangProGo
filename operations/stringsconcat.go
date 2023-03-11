package main

import "fmt"

func concat() {
	greeting := "Hello"
	language := "go"
	combinedString := greeting + ", " + language + "!"

	fmt.Println(combinedString)
	fmt.Println()
}
