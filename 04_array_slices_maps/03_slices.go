//if you're looping over an array, slice, map , string or reading from a channel you may use range.

package main

import "fmt"

func main() {
	slice1 := []float32{2.3, 3.4, 4.5}
	fmt.Println("length of the slice is ", len(slice1))
	fmt.Println("Index value of 0 and 2 indices are", slice1[0], slice1[2])

	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}
