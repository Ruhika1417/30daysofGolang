package main

import "fmt"

func foo() func() string {

	return func() string {
		return "I'm foo"
	}

}

func main() {

	//way 1
	counter := 1
	f := foo()
	fmt.Printf("%T\n", f)
	// f1 := f()
	// fmt.Println(f1, counter)
	// counter++

	//way 2
	f2 := foo()
	fmt.Println(f2(), counter)
	counter++

	//way 3
	fmt.Println(foo()(), counter)

}
