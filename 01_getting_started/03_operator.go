package main

import "fmt"

func main() {
	a := 10 //declaring and assigning
	fmt.Println("Value of a is:", a)
	a = 11 //only assigning (=) the updated value of a since we have already declared it.
	fmt.Println("Updated value of a is:", a)
	b := 10 + 20
	fmt.Println("Value of b is:", b)
}
