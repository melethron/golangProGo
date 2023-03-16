package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(e []Expense) (total float64) {
	for _, item := range e {
		total += item.getCost(true)
	}
	return
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

func main() {

	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"kayak", "watersports", 275},
			Service{"Boat Cover", 12, 89.50},
		},
	}

	for _, e := range account.expenses {
		fmt.Println("Expense:", e.getName(), "Price:", e.getCost(true))
	}
	fmt.Println("Print total:", calcTotal(account.expenses))
	fmt.Println()

	fmt.Println("Pointer method receivers")
	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = product
	product.price = 100
	fmt.Println("Product price:", product.price)
	fmt.Println("Expense price:", expense.getCost(false))

	fmt.Println("useing pointer")
	var expense2 Expense = &product
	product.price = 1000
	fmt.Println("Expense2 price:", expense2.getCost(false))
	fmt.Println()

	fmt.Println("Comparing interface values")
	var e1 Expense = &Product{name: "kayak"}
	var e2 Expense = &Product{name: "kayak"}
	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}

	fmt.Println("e1==e2:", e1 == e2)
	fmt.Println("e3==e4:", e3 == e4)
}
