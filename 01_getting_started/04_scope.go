/* whenever you want to declare variable outside of func body use var;
assigning w short declaration operator doesn't work outside of func body.
*/

package main

import "fmt"

var z = 10

var y int //if you don't assign value to a var zero value is assigned acc to the type. int:0 bool: false float:0.0

func main() {

	// fmt.Println("Value of x is:", x)       //won't work. out of scope.

	x := 10 //scope of x starts from here till the end of func body.
	fmt.Println("Value of x is:", x)

	call_y()

	fmt.Println("Value of z is:", z)

}

func call_y() {

	fmt.Println("Value of y is", y, "; Scope of y is entire program")

}
