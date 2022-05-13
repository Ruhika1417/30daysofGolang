// Create a program that uses a switch statement with the switch expression specified as a
// variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {

	favSport := "football"

	switch favSport {

	case "Cricket":
		fmt.Println("Cricket is my fav sport")

	case "football":
		fmt.Println("football is my fav sport ")

	case "Basketball":
		fmt.Println("Basketball is my fav sport")

	default:
		fmt.Println("ALl work and no play makes Jack a dull boy. JACK is me, I'm Jack!")

	}
}
