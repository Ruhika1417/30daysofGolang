//Create a program that uses a switch statement with no switch expression specified.
//without any expression it only excutes the true condition

//conclusion: without a switch expression only 1 true condition gets printed (FCFS).

package main

import "fmt"

func main() {
	x := 3
	switch {

	case !false:
		fmt.Println("true")

	case !true:
		fmt.Println("false")

	case x <= 3:
		fmt.Println("false case, won't be printed")

	case x >= 3:
		fmt.Println("true condition. x is greater than or equal to 3")

	}

}
