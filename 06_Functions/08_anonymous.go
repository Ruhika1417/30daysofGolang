package main

import "fmt"

var area = func(l int, b int) int {
	return l * b
}

func main() {

	//anonymous functions
	func(x int) {
		fmt.Println(x)

	}(40)

	//2
	area(10, 23)

}
