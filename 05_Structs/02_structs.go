//anonymous structs
//creating the structs without the struct types
//this is useful when you actually don't want to reuse the struct type
package main

import "fmt"

func main() {

	p1 := struct {
		name string
		age  int
		id   int
	}{
		name: "Ruhika",
		age:  20,
		id:   2222,
	}

	fmt.Println(p1)
	fmt.Println(p1.name)

}
