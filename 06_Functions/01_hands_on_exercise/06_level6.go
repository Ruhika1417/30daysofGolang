// Hands-on exercise #6
// ● Build and use an anonymous func
// ● Assign a func to a variable, then call that func

package main

import "fmt"

func apps() {
	fmt.Println("Spotify, YT, Pinterest")
}

func main() {

	a := func(a string) {
		fmt.Println(a)
	}

	a("It's me, I'm the anonymous function")

	b := apps
	b()

}
