// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
// TYPE []string which stores their favorite things. Store three records in your map. Print out all of
// the values, along with their index position in the slice.
// `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`
//add,delete a record to your map. Now print the map outusing the “range” loop

package main

import (
	"fmt"
)

func main() {
	//mistake: map[string]string{}

	m := map[string][]string{
		//`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},//probelm of redundancy. asked to remove []string

		`bond_james`: {`Shaken, not stirred`, `Martinis`, `Women`},

		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},

		`no_dr`: {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for i, v := range m {
		fmt.Println(i)
		for i2, v2 := range v {
			fmt.Println(i2, v2)
		}
	}

	m["Java"] = []string{"programming lang"}

	for i, v := range m {
		fmt.Print(i, v, "\n\n")
	}

	delete(m, "bond_james")
	for i, v := range m {
		fmt.Println(i, v)
	}

}
