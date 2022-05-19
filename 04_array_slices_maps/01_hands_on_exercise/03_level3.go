// hands-on exercise #3
// Using the code from the previous example, use SLICING to create the following new slices
// which are then printed:
// ● [42 43 44 45 46]
// ● [47 48 49 50 51]
// ● [44 45 46 47 48]
// ● [43 44 45 46 47]

package main

import "fmt"

func main() {
	s := [...]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Print(s, "\n\n")

	s1 := s[:5]
	fmt.Print(s1, "\n\n")

	s2 := s[5:]
	fmt.Print(s2, "\n\n")

	s3 := s[2:7]
	fmt.Print(s3, "\n\n")

	s4 := s[1:6]
	fmt.Println(s4)
	//fmt.Println(s[1:6])

}
