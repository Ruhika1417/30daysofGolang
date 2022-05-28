package main

import "fmt"

//method with a struct type receiver

type info struct {
	age int
}

type person struct {
	info
	name      string
	last_name string
}

func (p person) display() {
	fmt.Println(p.age, p.name)

}

func main() {

	d := person{info: info{32}, name: "Rachel", last_name: "Green"}

	fmt.Println(d)
	d.display()

}
