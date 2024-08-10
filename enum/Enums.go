package main

import "fmt"

type showStatus string 

const (
	Recieved showStatus = "Recieved"
	Confirmed = "Confirmed"
	Prepared = " Prepared"
	Delivered = "Deliverd"
)

func changeOrderStatus(status showStatus) {
	fmt.Println("Your current order status is ", status)
}

func main() {
	fmt.Println("Enums in golang")
	changeOrderStatus(Delivered)


}