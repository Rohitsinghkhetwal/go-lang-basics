package main

import (
	"fmt"
)

// we are creating the new order in other lang that how we create the constructor in oops

// type order struct {
// 	name   string
// 	status string
// 	price  float32
// 	qty    int
// }

// type product struct {
// 	name string
// 	dimension string
// 	color string
// 	qty int
// }

// func newProd (name string, dimension string, color string, qty int) *product{
// 	makeProd := product {
// 		name: name,
// 		dimension: dimension,
// 		color: color,
// 		qty: qty,
// 	}
// 	return &makeProd

// }

// func newOrder(name string, status string, price float32, qty int) *order {
// 	myOrder := order{
// 		name:   name,
// 		status: status,
// 		price:  price,
// 		qty:    qty,
// 	}
// 	return &myOrder
// }

type rectangle struct {
	len, width int
}

func (r * rectangle) area() int {
	return r.len * r.width


}

func main() {
	fmt.Println("Constructor in go lang")
	//result := newOrder("Shampoo", "Ordered", 455.5, 5)

	// res := newProd("maxx pro ultra", "2 * 2", "766", 4)
	// fmt.Println(res)
	// language := struct {
	// 	name string
	// 	status string
	// }{"Rohit", "Confirmed"}
	// fmt.Println(language)
	r := rectangle{len: 34, width: 56}
	fmt.Println("Area => ", r.area())
	



	

}
