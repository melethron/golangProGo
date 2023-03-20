package main

import (
	"fmt"
	"time"
)

func CalcStoretotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, group := range data {
		go group.totalPrice(category, channel)
	}
	for i := 0; i < len(data); i++ {
		storeTotal += <-channel
	}
	fmt.Println("total:", ToCurrency(storeTotal))
}

func (group ProductGroup) totalPrice(category string, resultChanel chan float64) {
	var total float64
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	resultChanel <- total
}
