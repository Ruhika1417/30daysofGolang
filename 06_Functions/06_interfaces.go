package main

import "fmt"

//interfaces are named collections of methods.

type rectangle interface {
	area() int
	perim() int
}

type value struct {
	length  int
	breadth int
}

func (v value) area() int {
	return v.length * v.breadth
}

func (v value) perim() int {
	return 2 * (v.length * v.breadth)
}

func main() {
	//accessing elements of the interface
	var r rectangle
	r = value{10, 12}
	fmt.Println("Area of rectangle", r.area())
	fmt.Println("Perim of rectangle", r.perim())

}
