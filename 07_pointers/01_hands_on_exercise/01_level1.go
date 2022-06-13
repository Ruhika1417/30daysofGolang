// Hands-on exercise #1
// ● Create a value and assign it to a variable.
// ● Print the address of that value.

package main

import "fmt"

func main() {

	b := 90
	var a *int = &b

	fmt.Println(a)

}
