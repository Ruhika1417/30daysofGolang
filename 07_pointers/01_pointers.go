package main

import "fmt"

func main() {
	a := 22
	fmt.Printf("type is %T and value is  %v\n", a, a)
	fmt.Printf("type is %T and value is %v\n", &a, &a)

	var b *int = &a
	fmt.Println(b, *b)
}
