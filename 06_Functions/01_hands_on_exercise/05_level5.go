/* Hands-on exercise #5
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 54
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle */

package main

import "fmt"

type square struct {
	length, width float32
}

type circle struct {
	radius int
}

func (s square) area() int {

	return int(s.length) * int(s.width)
}

func (c circle) area() int {
	pie := 3.14
	return 2 * int(pie) * c.radius
}

type shape interface {
	area() int
	// square()
	// circle()
}

func info(sh shape) {

	fmt.Println(sh.area())

}

func main() {

	s1 := square{
		length: 32.3,
		width:  23.4,
	}

	c1 := circle{
		radius: 69,
	}
	fmt.Println(s1.area())
	info(s1)
	info(c1)

}
