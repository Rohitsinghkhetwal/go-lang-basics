package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	prim() float32
}

type rectangle struct {
	h     float32
	width float32
}

type circle struct {
	radius float32
}

func (r rectangle) area() float32 {
	return r.h * r.width
}

func (r rectangle) prim() float32 {
	return 2*r.h + r.width

}

func (c circle) prim() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) area() float32 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("this is the value of -->", g)
	fmt.Println("This is the area of rect and circle", g.area())
	fmt.Println("This is the prim of rect and circle", g.prim())

}

type number struct {
	a int 
	b int
}

func (n number) add () int{
	return n.a + n.b
}

func (n number) subs()int {
	return n.a - n.b
}


func main() {
	fmt.Println("interface in golang")

	// res := number{a: 32 , b: 4}
	// fmt.Println("addition", res.add())
	// fmt.Println("subs", res.subs())

	c := rectangle{h: 34.2, width: 87.3}
	d := circle{radius: 3}

	measure(c)
	measure(d)

}
