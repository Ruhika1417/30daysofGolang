/* Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
solution: */

package main

import "fmt"

const (
	current_year = 2022                //index0 ??
	year1        = current_year + iota //index1
	year2        = current_year + iota
	year3        = current_year + iota
	year4        = current_year + iota
)

func main() {
	fmt.Println("The next 4 years are ", year1, year2, year3, year4)

}

//way2
// const (
// _ = iota
//	year1 = 2022 + iota
// )
