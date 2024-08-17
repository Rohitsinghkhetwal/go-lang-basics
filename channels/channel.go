package main

import (
	"fmt"
	"time"
)

func ProcessNum(result chan int) {
	fmt.Println("We are Processing the number to result => ", <-result)
}

func main() {
	fmt.Println("channels in go")
	result := make(chan int)
	go ProcessNum(result)
	result <- 44
	time.Sleep(time.Second * 2)

	
}