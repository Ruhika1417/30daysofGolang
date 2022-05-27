package main

import "fmt"

//multiple defers are executed in LIFO
func main() {

	defer first()
	second()
	defer third()
}

func first() {
	fmt.Println("I'm first")
}

func second() {
	fmt.Println("I'm second")
}

func third() {
	fmt.Println("I'm third")
}
