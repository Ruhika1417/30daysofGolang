package main

import "fmt"

func main() {
	slice1 := make([]int, 4, 5)
	fmt.Println("Length and capacity of the slice is ", len(slice1), ",", cap(slice1))
	fmt.Println("the slice is ", slice1)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = i
	}

	fmt.Println(slice1)
	fmt.Println("")

	slice1 = append(slice1, 4)
	fmt.Println("the slice is ", slice1)
	fmt.Println("Length and capacity of the slice after appending is ", len(slice1), ",", cap(slice1))
	fmt.Println("")

	//appending beyond the capacity
	slice1 = append(slice1, 5, 6)
	fmt.Println("the slice is ", slice1)
	fmt.Println("Length and capacity of the slice after append2 is ", len(slice1), ",", cap(slice1))
	fmt.Println("when extended beyong the capacity, the size of the underlyning array becomes double of the original capacity.")

}
