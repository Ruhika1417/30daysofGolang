package main

import "fmt"

func bar(a string) string {
	return a
}

func foo(f func(a string) string) {
	fmt.Println("string")

}

func main() {
	f := foo(bar("Ruhika"))
	fmt.Println(f)
}
