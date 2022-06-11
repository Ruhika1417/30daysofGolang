package main

import "fmt"

func main() {
	a := 22
	fmt.Printf("type is %T and value is  %v\n", a, a)
	fmt.Printf("type is %T and value is %v\n", &a, &a)

	// var b int = &a wrong coz it's 2 different types
	var b *int = &a
	fmt.Println(b) //b stores an address, with *b we're dereferencing that address
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 1111 //b stores an address, dereferencing the address and storing the value will change value of a
	fmt.Println("22 is changed to new value: ", a)

}
