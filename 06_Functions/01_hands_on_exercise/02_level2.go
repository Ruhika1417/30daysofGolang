/*Hands-on exercise #2
● create a func with the identifier foo that
○ takes in a variadic parameter of type int
○ pass in a value of type []int into your func (unfurl the []int)
○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
○ takes in a parameter of type []int
○ returns the sum of all values of type int passed in */

package main

import "fmt"

func foo(f ...int) int {
	total := 0
	for _, v := range f {
		total += v

	}
	return total

}

func bar(a []int) int {
	total := 0
	for _, v := range a {
		total += v

	}
	return total

}

func main() {

	x1 := []int{1, 2, 3, 4, 5, 9, 10}

	x2 := []int{111, 222, 333, 444, 555, 666, 777}

	f := foo(x1...)
	fmt.Println(f)

	b := bar(x2)
	fmt.Println(b)

}
