package main

import (
	"fmt"
	
)

// func ProcessNum(result chan int) {
// 	fmt.Println("We are Processing the number to result => ", <-result)
// }

// func worker(work chan bool) {
// 	fmt.Println("working .....")
// 	time.Sleep(time.Second * 3)
// 	fmt.Println("done !")
// 	work <- true
// }

func ping(pings chan <- string, msg string) {
	pings <- msg
}

func pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg

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

	// work := make(chan bool, 1)
	// go worker(work)
	// <-work	

	// channel direction 
	pings := make(chan string , 1)
	pongs := make(chan string, 1)
	ping(pings, "message passed")
	pong(pings, pongs)
	fmt.Println(<-pongs)





	
}