package main

import "fmt"

var a string = `ruhika` // for strings: " " or ` `
var b int               //zer

func main() {
	//string datatype
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	b = 31
	fmt.Printf("%%\n", b)
	fmt.Println("Value of b is", b)

}
