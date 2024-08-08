package main

import (
	"fmt"
	"time"
)

// structs in golang we can create our own data type or custom one
// we can create methods and data type of our own

type product struct {
	name      string
	price     float32
	qnty      int
	status    string
	createdAt time.Time // nano secondp presion

}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

type person struct {
	name string
	age int
	height float32
}

//recieving the order

func (p *product) changeStatus(status string) {
	p.status = status
}

func (o *order) changeAmount(amount float32) {
	o.amount = amount
}

func (p *product) changeName(name string) {
	p.name = name

}

// creating an object 

func newPerson(name string) *person{
	p := person{name: name}
	p.age = 34
	p.height = 4.6
	return &p
}

func main() {
	fmt.Println("Struct in go")

	var myProd = product{
		name:      "washing powder nirma",
		price:     43,
		qnty:      1,
		status:    "ordered",
		createdAt: time.Now(),
	}

	myProd.changeStatus("confirmed")

	fmt.Println("products are =>", myProd)

	order := order{
		id:        "45",
		amount:    435.4,
		status:    "orderd",
		createdAt: time.Now(),
	}

	order.changeAmount(600.43)
	myProd.changeName("TPhone 56 pro max")

	fmt.Println(myProd)
	//fmt.Println(order)
	per := person{"rohit", 24, 5.6}
	fmt.Println("This is the info about person", per)
	fmt.Println(newPerson("carl"))
	fmt.Println(myProd.name)


}
