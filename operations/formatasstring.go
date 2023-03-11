package main

import (
	"fmt"
	"strconv"
)

func formatAsString() {
	fmt.Println("Format values as strings")

	// Formatting bool
	fasVal1 := true
	fasVal2 := false
	str1 := strconv.FormatBool(fasVal1)
	str2 := strconv.FormatBool(fasVal2)
	fmt.Println("Formatted value1:", str1)
	fmt.Println("Formatted value2:", str2)

	// Formatting integers
	fasVal3 := 500
	base10Str := strconv.Itoa(fasVal3)
	base2Str := strconv.FormatInt(int64(fasVal3), 2)
	fmt.Println("Base10 string:", base10Str)
	fmt.Println("Base2 string:", base2Str)

	// Formatting floating point values
	fasVal4 := 70.95
	fString := strconv.FormatFloat(fasVal4, 'f', 2, 64)
	eString := strconv.FormatFloat(fasVal4, 'e', -1, 64)
	fmt.Println("Format F:", fString)
	fmt.Println("Format E:", eString)

	fmt.Println()
}
