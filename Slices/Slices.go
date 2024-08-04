package main

import "fmt"

func main() {
	fmt.Println("files of slices ")

	// var nums[]int
	// fmt.Println(nums == nil)

	// num := make([]int, 3)
	// fmt.Println(num)
	// num[0] = 43
	// num[1] = 42
	// num[2] = 56
	// fmt.Println(num)

	// var salaries[] int
	// fmt.Println(salaries)

	//this is an dynamic array that are used to declare a nil length of an array
	// make function is used to create a dynamic arrays

	// var noOfNames = make([]int, 0, 5);
	// fmt.Println(noOfNames)
	// noOfNames = append(noOfNames, 5)
	// noOfNames = append(noOfNames, 6)
	// noOfNames = append(noOfNames, 7)
	// noOfNames = append(noOfNames, 8)
	// noOfNames = append(noOfNames, 9)
	// noOfNames = append(noOfNames, 10)
	// noOfNames = append(noOfNames, 11)
	// noOfNames = append(noOfNames, 12)

	// fmt.Println(noOfNames)
	// fmt.Println(len(noOfNames))

	// var caps[] int
	// fmt.Println(caps)

	// var spcs = make([]int, 0, 5)
	// var num3 = make([]int, 0, 5)
	// spcs = append(spcs, 4)
	// num3 = make([]int, len(spcs))

	// copy(num3, spcs)
	// fmt.Println(spcs, num3)

	// slice operator
	// var nums = []int{2,3,4,5,6}

	// fmt.Println(nums[0:3])

	//declare a slice
	// var nums4 = []int {4,6,7,8,9}
	// fmt.Println(nums4[2:3])
	// var ytu[]string
	// ytu = append(ytu, "Rohit singh Khetwal")
	// fmt.Println(ytu)

	var open = []int{5,8,2,3,0,8,9}
	fmt.Println(open[:3])
	fmt.Println(open[3:])

	// 2d array

	maxi := [2][3]int{{2,3,4}, {8,4,56}}
	fmt.Println(maxi)





}
