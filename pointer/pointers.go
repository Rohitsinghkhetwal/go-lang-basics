package main

import "fmt"

func changeByRefrence(marks *int) {
	*marks = 43
	//fmt.Println("marks inside the marks", *marks)
	
}

func changeByValue(num int) {
	num = 44
	//fmt.Println("this is inside the value func", num)
}

// passin the address of slices in functions

func passSlices(num *[]int) {
	*num = append(*num, 9345)
	fmt.Println("The prices of stocks are =>", *num)

}

func main() {
	fmt.Println("Pointers in go lang")
	marks := 56
	changeByRefrence(&marks)
	changeByValue(marks)

	stockPrices := []int{34,54,32,25,65}
	passSlices(&stockPrices)

	fmt.Println("values of stockprices are =>", stockPrices)


	
}