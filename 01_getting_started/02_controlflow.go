package main

import "fmt"

func main() {

	n, _ := fmt.Println(" A program to show the control flow ")
	fmt.Println(n) //bytes

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	secondfunc()

}

func secondfunc() {
	fmt.Println("I'm the second func")
}
