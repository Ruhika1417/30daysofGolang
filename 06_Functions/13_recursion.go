package main

import "fmt"

func main() {

	f := factorial(3)
	fmt.Println(f)

	//2

	fib(7)

}

func factorial(n int) int {
	for n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//var fib func(n int) int

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//code for factorial using loop

// func factorial(n int) int {
// 	a := 1
// 	// or: for ; n>0; n--
// 	for i := 1; i <= n; i++ {
// 		a *=i
// 	}
// 	return a
// }
