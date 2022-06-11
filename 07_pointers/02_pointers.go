package main

import "fmt"

func main() {
	var a int = 32
	pointers(&a)
	fmt.Println("a value", a) //34
}

func pointers(b *int) {
	fmt.Println("b before", b)
	fmt.Println("b before", *b)
	*b = 34 //dereferencing b and changing the value
	fmt.Println("b after ", *b)
	fmt.Println("b after ", b)

}
