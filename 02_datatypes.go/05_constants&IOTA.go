package main

import (
	"fmt"
)

// way 1
const a = 10

//way 2 // untyped const
const (
	b = iota
	c = iota
	_ // to skip value.
	d = iota
)
 fmt.Println(" ") 
 
// typed const. //iota is reset to 0 when the const word appears again.
const (
	e float32 = iota
	f float32 = iota
)

func main() {

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n%T %v\n%T %v\n", b, b, c, c, d, d)
	fmt.Printf("%T %v\n", e, e)
	fmt.Printf("%T %v\n", f, f)

}
