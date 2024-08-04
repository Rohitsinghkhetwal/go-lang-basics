package main

import (
	"fmt"
	"maps"
)


func main() {
	fmt.Println("Map function")

	//declaration of a map to store a unique numbers and string	

	uniqueNum := make(map[string]int)
	uniqueNum["medium"] = 30
	uniqueNum["maximum"] = 78

	fmt.Println(uniqueNum["medium"])
	// for declaring the nums without intialization

	nums := make(map[int]int)
	nums[0] = 1
	nums[2] = 4

	fmt.Println(nums[0])
	delete(nums, 0)
	fmt.Println(nums[0])
	clear(nums)
	fmt.Println(nums[0])
	// another method of declaring the map 
	//declaring the the map with initialization with more than one value
	 n := map[string] int {"rohit" : 1, "bar": 3}
	 fmt.Println(n)
	 n2 := map[string]int{"rohit" : 2, "bar": 3}
	 // we can compare each other as well

	 if maps.Equal(n, n2) {
		fmt.Println("They are both equal to each !")
	 }else {
		fmt.Println("how's the josh !")
	 }

}