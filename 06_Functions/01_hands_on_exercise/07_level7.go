// Hands-on exercise #8
// ● Create a func which returns a func
// ● assign the returned func to a variable
// ● call the returned fun

package main

import "fmt"

func sum() func(a, b int) int {

	return func(a, b int) int {
		return (a + b)

	}
}

func main() {

	a := sum()(3, 4)
	fmt.Println(a)

}
