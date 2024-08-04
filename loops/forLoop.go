package main

import "fmt"

func main() {
	// while loop fashion
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for loop fashion

	// for i := 0; i < 20; i++ {
	// 	fmt.Println(i)
	// }

	// third fashion

	// for i := range 6 {
	// 	if i%2 == 0 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	var price int = 23
	fmt.Println(price)

	for i := range 10 {
		fmt.Println(i)
	}

	myname := "Rohit singh"
	fmt.Println(myname)

	const group string = "mrwhostheboss"
	fmt.Println(group)

	var name, fullname, fathername string = "rohit", "singh", "khetwal"
	fmt.Println(name, fullname, fathername)

	// while loop
	i := 0
	for i <= 9 {
		fmt.Println(i)
		i++
	}

}
