package main

import "fmt"

func main() {
	xii := []int{1, 2, 3, 4, 5, 6}
	//a := sum(xii...)
	//fmt.Println("all numbers", a)

	a2 := even(sum, xii...)
	fmt.Println(a2)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total

}

func even(f func(x ...int) int, vi ...int) int {

	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}

	return f(y...)
}
