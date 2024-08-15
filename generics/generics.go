package main

import (
	"fmt"
)

// basically generics in go lang are the way to avoid DRY functions it's syntax is given below

// func strSlice[G int | string | bool](nums[]G ) {
// 	for _, nums := range nums {
// 		fmt.Println(nums)
// 	}
// }

// func strSlice(names[] string) {
// 	for _, names := range names {
// 		fmt.Println("favrouite lang are =>", names)
// 	}
// }

// we can use this on struct 

type stack[T any] struct {
	elements []T

}

// In this way we can use generics in go with slices and structs 

type person [T int | string]struct {
	name string
	ids  []T
	isActive bool
	maxCGPA float32
}


func main() {
// 	fmt.Println("Generics in golang !")
// 	nums := []int{3,4,5,6}
// 	nums2 := []string{"rohit", "max", "chow"}
// 	strSlice([]string{"golang", "javascript", "nodejs", "java"})
// 	bl := []bool{true, false, true}
// 	strSlice(bl)
   myStack := stack[string]{
		elements: []string{"golang"},
	 }
	 fmt.Println(myStack)

	 newPerson := person[int] {
		name: "Rohit",
		ids: []int{1,2,3,4},
		isActive: true,
		maxCGPA: 8.5,
	 }

	 fmt.Println(newPerson)
  }