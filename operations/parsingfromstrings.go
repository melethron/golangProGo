package main

import (
	"fmt"
	"strconv"
)

func stringConv() {
	fmt.Println("Parsing from strings")
	val1 := "true"
	val2 := "false"
	val3 := "not true"

	if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
		fmt.Println("Bool1", bool1)
	} else {
		fmt.Println("cannot parse", val1)
	}

	if bool2, b2err := strconv.ParseBool(val2); b2err == nil {
		fmt.Println("Bool1", bool2)
	} else {
		fmt.Println("cannot parse", val2)
	}
	if bool3, b3err := strconv.ParseBool(val3); b3err == nil {
		fmt.Println("Bool3", bool3)
	} else {
		fmt.Println("Cannot parse", val3)
	}

	fmt.Println()
}
