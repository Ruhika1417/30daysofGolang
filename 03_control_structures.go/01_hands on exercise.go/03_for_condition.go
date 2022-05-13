// Create a for loop using this syntax for condition { }
//Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	age := 20
	fmt.Println("years that I have been alive:")
	for i := 1; i <= age; i++ {

		fmt.Printf("%d\n", i)

	}
}
