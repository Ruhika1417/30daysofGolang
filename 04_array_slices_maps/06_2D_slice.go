package main

import "fmt"

func main() {

	s1 := []string{"Coffee", "Butter"}
	fmt.Println(s1)

	s2 := []string{"Ice cream", "Scotch"}
	fmt.Println(s2)

	s3 := [][]string{s1, s2}
	fmt.Println(s3)

	//way 2
	s4 := [][]string{
		{"A", "B"},
		{"C", "D"},
	}
	fmt.Println(s4)
}
