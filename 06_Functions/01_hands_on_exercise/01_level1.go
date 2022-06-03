// Hands on exercise 1
// ○ create a func with the identifier foo that returns an int
// ○ create a func with the identifier bar that returns an int and a string
// ○ call both funcs
// ○ print out their results

package main

import "fmt"

func foo(a int) int {

	return a
}

func bar(a int, s string) (int, string) {

	return a, s
}

func main() {
	f := foo(32)
	fmt.Println(f())

	b1, b2 := bar(32, "Bar")
	fmt.Println(b1, b2)

}
