//closure helps us limit the scope of variables
//anonymous functions can form a closure

package main

import "fmt"

func main() {

	//1
	var a int = 10

	func() {
		for i := 0; i < 5; i++ {
			a += i
		}
		fmt.Println(a)
	}()

	//2

	i1 := counter()
	fmt.Println(i1())
	fmt.Println(i1())
	i2 := counter()
	fmt.Println(i2())

}

func counter() func() int {
	var b int
	return func() int {
		b++
		return b
	}
}
