// hands on exercise 9
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// ● pass a func into a func as an argument.

package main

import "fmt"

func main() {

	//anony func

	a := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]

	}

	x := callback(a, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)

}

func callback(f func(xi []int) int, xiii []int) int {

	n := f(xiii)
	fmt.Println(n)
	n++
	fmt.Println(n)

	return n

}
