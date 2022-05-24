// Hands-on exercise #4
// Create and use an anonymous struct.
//soln:anonymous structs do not have any struct type.

package main

import "fmt"

func main() {

	gadgets := struct {
		laptop string
		phone  string
		watch  int
	}{
		laptop: "DELL",
		phone:  "IPhone",
		watch:  1,
	}
	fmt.Println(gadgets)

}
