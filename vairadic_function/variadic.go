package main

import "fmt"

func sum(max  ...int)int {
	total := 0

	for _, max := range max {
		total = total + max
	}
	return total
}

func main() {
	fmt.Println("variadic function in go")

	// we can pass a slice also in this 
	var nums = []int{3,5,6,7}
	fmt.Println(nums)

	result := sum(nums...)
	fmt.Println(result)

}