package main

import (
	"fmt"
	"sort"
)

func slices() {
	fmt.Println("Slices")
	sliceNames := make([]string, 3)
	sliceNames[0] = "Kayak"
	sliceNames[1] = "Lifejacket"
	sliceNames[2] = "Paddle"

	fmt.Println(sliceNames)

	sliceNames2 := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(sliceNames2)

	sliceNames2 = append(sliceNames2, "Hat", "Gloves")
	fmt.Println(sliceNames2)

	appendedNames := append(sliceNames, "Hat", "Gloves")
	sliceNames[0] = "Canoe"
	fmt.Println("sliceNames:", sliceNames)
	fmt.Println("appendedNames:", appendedNames)
	fmt.Println()

	// Allocating additional capacity
	fmt.Println("Allocating additional capacity")
	sliceNames3 := make([]string, 3, 6)
	sliceNames3[0] = "Kayak"
	sliceNames3[1] = "Lifejacket"
	sliceNames3[2] = "Paddle"
	fmt.Println("len:", len(sliceNames3))
	fmt.Println("cap:", cap(sliceNames3))
	appendedNames2 := append(sliceNames3, "Hat", "Gloves")
	sliceNames3[0] = "Canoe"
	fmt.Println("names:", sliceNames3)
	fmt.Println("appendedNames2:", appendedNames2)
	fmt.Println()

	// Append one slice to another
	fmt.Println("appending slices")
	sliceNames4 := make([]string, 3, 6)
	sliceNames4[0] = "Kayak"
	sliceNames4[1] = "Lifejacket"
	sliceNames4[2] = "Paddle"
	moreNames := []string{"Hat", "Gloves"}
	appendedNames3 := append(sliceNames4, moreNames...)
	fmt.Println("appended slices:", appendedNames3)
	fmt.Println()

	// Creating slices from existing arrays
	fmt.Println("Creating slices from existed arrays")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("Somenames:", someNames)
	fmt.Println("Somenames len:", len(someNames), "cap", cap(someNames))
	fmt.Println("All names:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))

	someNames = append(someNames, "Gloves")
	someNames = append(someNames, "Boots")
	fmt.Println("Somenames:", someNames)
	fmt.Println("Somenames len:", len(someNames), "cap", cap(someNames))
	fmt.Println("All names:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	fmt.Println()

	// Specifying capacity, when creating slice
	_ = products[1:3:3]

	// Creating slices from other slices
	fmt.Println("Creating slices from other slices")
	products = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames2 := products[1:]
	someNames2 := allNames2[1:3]
	allNames2 = append(allNames2, "Gloves")
	allNames2[1] = "Canoe"
	fmt.Println("Somenames:", someNames2)
	fmt.Println("All names:", allNames2)
	fmt.Println()

	// Copy function
	fmt.Println("Copy function")
	products = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames3 := products[1:]
	someNames3 := make([]string, 2)
	copy(someNames3, allNames3)
	fmt.Println("Allnames:", allNames3)
	fmt.Println("Somenames:", someNames3)
	fmt.Println()

	// Uninitialized slice pitfall
	fmt.Println("\"Uninitialized Slice pitfall\"")
	var uninitialized []string
	copy(uninitialized, allNames3)
	fmt.Println("Uninitialized slice:", uninitialized)
	fmt.Println()

	// Specifying ranges, when copying slices
	fmt.Println("Ranges, when copying slices")
	products = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames4 := products[1:]
	someNames4 := []string{"Boots", "Canoe"}
	copy(someNames4[1:], allNames4[2:3])
	fmt.Println("Allnames:", allNames4)
	fmt.Println("Somenames:", someNames4)
	fmt.Println()

	// Copying slices with different sizes
	fmt.Println("Copying slices with different slices")
	products2 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts2 := []string{"Canoe", "Boots"}
	copy(products2, replacementProducts2)
	fmt.Println("Products:", products2)
	products3 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts3 := []string{"Canoe", "Boots"}
	copy(products3[0:1], replacementProducts3)
	fmt.Println("Products:", products3)
	fmt.Println()

	// Deleting slice elements
	fmt.Println("Delete elements from slice")
	products4 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	deleted := append(products4[:2], products4[3:]...)
	fmt.Println("Deleted:", deleted)
	fmt.Println()

	// Sorting slices
	fmt.Println("Sorting slices")
	products5 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	fmt.Println("Before sorting:", products5)
	sort.Strings(products5)
	fmt.Println("After sorting:", products5)
	fmt.Println()

	// Getting the array, underlying the slice
	fmt.Println("Getting underlying array")
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr
	fmt.Println("Underlying array is", array)
	fmt.Println()

}
