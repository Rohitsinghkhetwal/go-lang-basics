package main

import "fmt"

func main() {
	fmt.Println("Range in go ")

	nums := []int{2, 3, 4, 5}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	for i, nums := range nums {
		fmt.Println(nums, i)
	}

	arr := map[string] string{"fname": "Rohit", "Lastname": "Khetwal", "roll": "s. eng."}

	for k, v := range arr {
		fmt.Println(k, v)
	}
	nums2 := make(map[int]int)
	nums2[3] = 54
	nums2[4] = 87
	fmt.Println(nums2)

}
