// Hands on exercise 3
// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

package main

import "fmt"

func deferencing() {
	fmt.Println("I will only run after the func containing me exists")
}

func normal() {
	fmt.Println("I will execute as soon as I'm called")
}

func main() {

	defer deferencing()
	normal()

}
