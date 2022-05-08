// kb = 1000000 // mb = 10000000000000(10 extra zero's)
package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10) //shift 1 through 10 digits on the left. here iota =1
	mb = 1 << (iota * 10) // shift 1 throu 20 digits on the left
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%v\t\t %b\n", kb, kb) //value in decimal and binary form.
	fmt.Printf("%v\t\t %b\n", mb, mb)
	fmt.Printf("%v\t %b\n", gb, gb)
}
