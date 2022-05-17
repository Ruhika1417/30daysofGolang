package main

import "fmt"

func main() {

	slice1 := []int{2, 3, 45, 46}
	slice1 = append(slice1, 45)
	fmt.Println(slice1)

	//appending one slice to another slice

	slice2 := []int{47, 48}
	slice3 := append(slice1, slice2...) //dots afterwards means take all the values of slice2 and put them right here.
	fmt.Println("The new slice is :", slice3)

}
