package main

import "fmt"



func main() {
	fmt.Println("Arrays in go lang")
	//initialization a an arrays and putting in
	// var nums[4] int 
	// fmt.Println(nums)

	// nums[3] = 10
	// fmt.Println(nums)

	// Arr := [5]int{1,2,3,4,5}
	// fmt.Println(Arr)

	// Arr = [...]int{3,4,5,6,7}
	// fmt.Println(Arr)

	// prices := [6]int{23,54,34,22,98,25}

	// for i := 0; i <= 5; i++ {
	// 	for j := i + 1; j <= 5; j++ {
	// 		if prices[i] > prices[j] {
	// 			temp := prices[i]
	// 			prices[i] = prices[j]
	// 			prices[j] = temp
	// 		}

	// 	}
	// }
	// fmt.Println(prices)

	var nums[4]int 
	fmt.Println(len(nums))

	nums[0] = 67
	fmt.Println(nums)

	books := [5]int{2,3,4,5,6}
	fmt.Println(books)

	//2d arrays
	maxArr := [2][5]int{{3,4,5,6,7}, {54,65,76,4323,23}}
	fmt.Println(maxArr)

	

}
