//a for statement iteration can be controlled by 1. a single statement, 2. a for clause 3. a range clause ; 4 just for { } also works.

package main

import "fmt"

func main() {
	//way 1
	a, b := 110, 1198
	for a < b {
		fmt.Printf("Value of a: %v is less value of b: %v. Hence b is the greater number\n", a, b)
		break

	}

	//way 2.
	// Syntax for init; condition; post { }
	//here nested loop.
	for i := 1; i < 3; i++ {
		fmt.Println("I'm the outer loop: ", i)
		for j := 1; j < 3; j++ {
			fmt.Println("\t\tI'm the inner loop", j)

		}
	}

	//way 3

	array1 := [2]int{11, 12}

	for i, v := range array1 {
		fmt.Println("Value and index of the array are", i, v)
	}

	// to omit the second value use this
	for i := range array1 {
		fmt.Println("Index of the array is ", i)

	}
	// to omit the first return value
	for _, v := range array1 {
		fmt.Println("Value of the array is ", v)
	}

	// way 4

	for {
		c := a * b
		fmt.Println(c)
		break

	}

}
