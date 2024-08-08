package main

import "fmt"

// we are creating the new order in other lang that how we create the constructor in oops

type order struct {
	name   string
	status string
	price  float32
	qty    int
}

func newOrder(name string, status string, price float32, qty int) *order {
	myOrder := order{
		name:   name,
		status: status,
		price:  price,
		qty:    qty,
	}
	return &myOrder
}

func main() {
	fmt.Println("Constructor in go lang")
	result := newOrder("Shampoo", "Ordered", 455.5, 5)

	fmt.Println(result)

}
