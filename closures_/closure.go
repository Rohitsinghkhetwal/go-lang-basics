package main

import "fmt"

//closure in go is a function that returns a function 
//this is type of features that are inherited from javascript 
// it has access to the outer execution context


func intSeq() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

func main() {
	fmt.Println("closures in go")


	nextSeq := intSeq()
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
}