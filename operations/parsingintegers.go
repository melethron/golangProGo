package main

import (
	"fmt"
	"strconv"
)

func parsingIntegers() {
	fmt.Println("Parsing Integers")
	piVal1 := "500"

	// Parse integer
	if int1, int1err := strconv.ParseInt(piVal1, 10, 10); int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse", piVal1)
	}

	// Parse binary
	fmt.Println("Parsing binary")
	bVal1 := "10010"
	if int2, int2err := strconv.ParseInt(bVal1, 2, 0); int2err == nil {
		fmt.Println("Parsed value:", int2)
	} else {
		fmt.Println("Cannot parse:", bVal1)
	}

	// Parse binary with 0b prefix and let ParseInt to decide the base
	bVal2 := "0b1001"
	if int3, int3err := strconv.ParseInt(bVal2, 0, 8); int3err == nil {
		fmt.Println("Parsed value:", int3)
	} else {
		fmt.Println("Cannot parse", bVal2)
	}

	// Atoi function
	piVal4 := "100"
	int4, int4err := strconv.Atoi(piVal4)
	if int4err != nil {
		fmt.Println("Cannot parse", piVal4, int4err)
	} else {
		fmt.Println("Parsed value:", int4)
	}

	// Parsing float
	piVal5 := "70.96"
	float1, float1err := strconv.ParseFloat(piVal5, 64)
	if float1err != nil {
		fmt.Println("Cannot parse", piVal5, float1err)
	} else {
		fmt.Println("Parsed value:", float1)
	}

	piVal6 := "4.895e+01"
	float2, float2err := strconv.ParseFloat(piVal6, 64)
	if float2err != nil {
		fmt.Println("Cannot parse", piVal6, float2err)
	} else {
		fmt.Println("Parsed value:", float2)
	}

	fmt.Println()
}
