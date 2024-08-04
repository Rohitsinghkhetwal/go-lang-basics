package main

import "fmt"

func TraverseArr(arr []int) {
	// function to sort the slice i go
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}

	}
}

func main() {
	fmt.Println("|functions in go |")
	// a := 32
	// b := 35

	// var res = sum(a, b)
	// fmt.Println(res)

	// c, d := vals()
	// fmt.Println(c)
	// fmt.Println(d)

	// _, e := vals()
	// fmt.Println("this is a emty func",e)

	nums := []int{3, 6, 4, 9, 8, 7, 2}
	TraverseArr(nums)
	fmt.Println(nums)
	// this step to reverse the array

	start := 0
	end := len(nums) - 1

	for start < end {
		if nums[start] < nums[end] {
			temp := nums[start]
			nums[start] = nums[end]
			nums[end] = temp
			start++
			end--
		}
	}

	fmt.Println(nums)


}
