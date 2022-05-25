// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors.
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.

package main

import "fmt"

func main() {

	//mistake var person struct{}

	type person struct {
		first_name string
		last_name  string
		icecream   []string
	}

	p1 := person{
		first_name: "Ruhika",
		last_name:  "B",
		icecream: []string{
			"coffee",
			"coconut"},
	}

	p2 := person{
		first_name: "Megha",
		last_name:  "B",
		icecream: []string{
			"strawberry", "chococream"},
	}

	fmt.Println(p1.first_name)

	fmt.Println(p1.last_name)
	for _, v := range p1.icecream {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", p1.icecream)

	fmt.Println("Icecream flavours of person2:")
	for _, v := range p2.icecream {
		fmt.Println(v)
	}

}
