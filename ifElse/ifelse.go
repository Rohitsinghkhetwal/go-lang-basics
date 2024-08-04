package main

import "fmt"

func main() {
	fmt.Println("If else statements")
	if 5 % 2 == 1 {
		fmt.Println("Number is odd")
	}else {
		fmt.Println("Number is even")
	}

	if num := 9; num < 0 {
		fmt.Println("Number is negative")
	}else if num < 9 {
		fmt.Println("Number is one digit ")
	}else {
		fmt.Println("Number is greater than 9")
	}
}