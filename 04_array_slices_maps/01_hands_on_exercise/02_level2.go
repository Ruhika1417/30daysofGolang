// hands on exercise #2
// ● Using a COMPOSITE LITERAL:
// ● create a SLICE of TYPE int
// ● assign 10 VALUES
// ● Range over the slice and print the values out.
// ● Using format printing
// ○ print out the TYPE of the slice.

package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", s)
}
