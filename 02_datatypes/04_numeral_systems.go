package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := "256"
	fmt.Printf("%T %v\n", a, a)

	// converting string to int

	bs := []byte(a)
	fmt.Println(bs)

	b, _ := strconv.Atoi(a)

	fmt.Printf("%b is the value in binary\n", b)
	fmt.Printf("%d,is the value in decimal\n", b)
	fmt.Printf("%#x is the value in hexadecimal\n", b)
	fmt.Printf("% X is the value for each byte w base 2 uppercase", a)

	fmt.Println(" ")

	// can do for each byte as well.

	index0 := bs[0]
	fmt.Println(index0)
	fmt.Printf("%b is the value in binary\n", index0)
	fmt.Printf("%d,is the value in decimal\n", index0)
	fmt.Printf("%#x is the value in hexadecimal\n", index0)

}
