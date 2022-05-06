package main

import "fmt"

type number int

var x number

func main() {
	fmt.Printf("Value of x is %v\nType of x is %T\n", x, x)
	//fmt.Printf("Type of x is %T\n", x)
	x = 42
	fmt.Printf("New value of x is %v", x)

}
