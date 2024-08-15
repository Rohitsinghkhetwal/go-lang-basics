package main

import (
	"fmt" 
	"time"
)

// what is the purpose of goroutines ? basically it is a lightweight thread of execution 
// that runs two different program in parallel it is very fast 
// some of examples are here !

func Tasks(someone string) {

	fmt.Println("doing task here !")
	for i:= 0; i <=3; i++ {
		fmt.Println(someone, ":", i)
	}
}

func main() {
	fmt.Println("go Routine in golang")
	Tasks("direct")
	go Tasks("go routines")


	go func(message string) {
		fmt.Println(message)

	}("going")

	time.Sleep(time.Second * 2)
	fmt.Println("Done Dona Done !")

}