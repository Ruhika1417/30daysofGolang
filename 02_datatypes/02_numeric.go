//uint, int, float, complex, byte, rune are the numeric types.

package main

import "fmt"

var a uint = 10
var b int = 67
var c float32 = 67.6777
var d complex64 = 12.7 + 10i //i is underoot of -1.
var e byte = 255             // byte is alias of unit 8
var f rune = -1899999999     //rune is alias of int32

func main() {
	fmt.Printf("%v %T\n%v %T\n%v\n%T %v\n", a, a, b, b, c, d, d)
	fmt.Printf("%v %T\n%v %T\n", e, e, f, f)

}
