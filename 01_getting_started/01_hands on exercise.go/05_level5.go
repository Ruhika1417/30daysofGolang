/* Continuing level 4
b. now do this
i. now use CONVERSION to convert the TYPE of the VALUE stored in “x”
to the UNDERLYING TYPE
1. then use the “=” operator to ASSIGN that value to “y”
2. print out the value stored in “y”
3. print out the type of “y */

package main

import "fmt"

type number int

var x number
var y int

func main() {
	fmt.Printf("Value of x is %v\nType of x is %T\n", x, x)
	//fmt.Printf("Type of x is %T\n", x)
	x = 42
	fmt.Printf("New value of x is %v", x)
	y = int(x)
	fmt.Println("Value of y is ", y)

}
