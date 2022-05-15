package main

import "fmt"

func main() {

	if !true {
		fmt.Println("false")

	}

	if !false {
		fmt.Println("true")

	}

	//you can have multiple if statements. No need of parenthesis around condition in if/else.

	if x := 1; x == 1 {
		fmt.Println(x, "equal to 1")

	}
	//fmt.Println(x)// out of scope

	if x := 1; x > 2 {
		fmt.Println(x, "not equal to 1")

	}

	if x := 1; x < 1 {
		fmt.Println(x, " not equal to 1")
	}

	//way 2
	x := 2 //increaing the scope of x till }
	if x == 2 {
		fmt.Print("Equal values: ")
	}
	fmt.Println(x)

}
