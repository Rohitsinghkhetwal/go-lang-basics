package main

import (
	"fmt"
	"time"
)

// func ProcessNum(result chan int) {
// 	fmt.Println("We are Processing the number to result => ", <-result)
// }

func worker(work chan bool) {
	fmt.Println("working .....")
	time.Sleep(time.Second * 3)
	fmt.Println("done !")
	work <- true
}

func main() {
	// fmt.Println("channels in go")
	// result := make(chan int)
	// go ProcessNum(result)
	// result <- 44
	// time.Sleep(time.Second * 2)



	// channel buffering

	// message := make(chan string , 2)

	// message <- "Ready for buffer"
	// message <- "Done !"

	// fmt.Println(<-message)
	// fmt.Println(<-message)

	// channel synchronization

	work := make(chan bool, 1)
	go worker(work)
	<-work	




	
}