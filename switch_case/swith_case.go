package main

import "fmt"

func main() {
	fmt.Println("Hey this is a switch case")

	nums := 5
	fmt.Println("nums is written as", nums, " as")
	switch nums {
	case 1: 
	     fmt.Println("Number is one")
	case 2: 
	    fmt.Println("Number is two")
	case 3:
		  fmt.Println("Number is three")
	default:
		fmt.Println("I won't see anyhting i Dont't know counting")
	}
}