/* Write a program that
● assigns an int to a variable
● prints that int in decimal, binary, and hex
● shifts the bits of that int over 1 position to the left, and assigns that to a variable
● prints that variable in decimal, binary, and hex */

package main

import "fmt"

var a int = 880

func main() {
	fmt.Printf(" Value of a in decimal is %d in binary is %b and in hex is %x\n", a, a, a)

	b := a << 1
	fmt.Printf(" Value of b in decimal is %d in binary is %b and in hex is %x\n", b, b, b)

	c := 1 << 1
	fmt.Printf(" Value of a in decimal is %d in binary is %b and in hex is %x\n", c, c, c)

}
