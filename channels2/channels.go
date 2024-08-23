package main

import (
	"fmt"

)

// func sum(result chan int , num2 int, num3 int) {
// 	res := num2 + num3
// 	result <- res
// }

// goroutine sychronizer

// func task(done chan bool ) {
// 	defer func() {done <- true}()
// 	fmt.Println("Data processing ...")

// }
// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() {done <- true}()

// 	for em:= range emailChan {
// 		fmt.Println("sending email to ",em)
// 		time.Sleep(time.Second )
// 	}

// }


func main() {
	fmt.Println("channels in go lang")
	// result := make(chan int )
	// // go routines
	// go sum(result, 5, 5)
	// trip := <- result
	// fmt.Println(trip)

	// done := make(chan bool )
	// go task(done)
	// <- done

	// buffered channel
	// emailChan := make(chan string, 100)
	// done := make(chan bool )
	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("Done !!!!")
	// close(emailChan)
	// <- done(

	// handle multiple channels

	chan1 := make(chan int )
	chan2 := make(chan string)

	// inline func with closure

	go func() {
		chan1 <- 45
	}()

	go func() {
		chan2 <- "ROhit singh"
	}()

	// iterate through

	for i:= 0; i < 2; i++ {
		select {
		case IntVal := <- chan1:
			fmt.Println("Recieved data from IntChannel", IntVal)
		case StrVal := <- chan2:
			fmt.Println("Recieved data from StrChannel", StrVal)
		}
	}





}