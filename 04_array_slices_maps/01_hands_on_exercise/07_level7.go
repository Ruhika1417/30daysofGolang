// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
// slice:
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {

	s := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
