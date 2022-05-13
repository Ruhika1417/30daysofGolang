// Create a for loop using this syntax ** for condition { }**
//Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	born_y := 2001

	fmt.Println("years that I have been alive:")
	for born_y <= 2022 {
		fmt.Println(born_y)
		born_y++
	}

}

//works but use 1 condition only , as per the task.

// for i := 1; i <= age; i++ {

// 	fmt.Printf("%d\n", i)

// }
