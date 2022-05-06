package main

import "fmt"

//creating our own type "car".
type car int

var b int

func main() {

	var a car = 10

	fmt.Sprintln(a)

	fmt.Println("%T", a)

	fmt.Printf("%T", a)
	//b = a
	fmt.Println(b)
	b = int(a)
	fmt.Println(b)

	fmt.Printf("%#v", a)
}
