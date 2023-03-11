package main

import "fmt"

func comparison() {
	fmt.Println("Comparison")
	first := 100
	const second = 200

	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second

	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
	fmt.Println()

	// Comparing pointers
	fmt.Println("Comparing porinters")
	pFirst := 100
	pSecond := &pFirst
	pThird := &pFirst

	alpha := 100
	beta := &alpha
	//Compare address locations
	fmt.Println(pSecond == pThird)
	fmt.Println(pSecond == beta)
	//Compare values

	fmt.Println(*pSecond == *pThird)
	fmt.Println(*pSecond == *beta)
	fmt.Println()
}
