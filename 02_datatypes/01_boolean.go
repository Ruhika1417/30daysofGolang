package main

import "fmt"

var a bool

func main() {

	fmt.Println(a, "is the zero value of bool")
	a = true
	fmt.Println(a, "is the updated value of bool")

	b := 34
	c := 97
	fmt.Println(b != c)

}
