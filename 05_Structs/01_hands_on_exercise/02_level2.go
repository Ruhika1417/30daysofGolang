// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.

package main

import "fmt"

func main() {

	//mistake var person struct{}

	type person struct {
		first_name string
		last_name po
		icecream   string
	}

	// p0:= map[string][string]{	
	// 	first_name:  "B1",
	// 	last_name:  "B",
	// 	icecream:   "Coffee",	
	// }

    p1 := person{
		first_name: "Ruhika",
		last_name:  map[string]string{"Ruhika", 67,90}
		icecream:   "Coffee",
	}

	p2 := person{
		first_name: "Megha",
		last_name:  "B",
		icecream:   "Coconut",
	}

	fmt.Println(p1.icecream)
	fmt.Println(p2.icecream)

	for i, v:= range p0{
		fmt.Println(i,v)
	}

}
