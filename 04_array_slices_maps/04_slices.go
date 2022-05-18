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

	// **special case appending string to a byte**
	//bs1 := []byte("hello")
	//bs2 := append(bs1, "hello"...)
	//fmt.Println(bs2)
	slice45 := append([]byte("hello"), "world"...)
	fmt.Println(slice45)

	//deleting elements of a slice using append
	slice4 := append(slice1[:1], slice2[:]...)
	fmt.Println(slice4)

}
