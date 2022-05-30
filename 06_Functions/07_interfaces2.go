package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	fname string
	lname string
}

type info struct {
	person
	age int
	ltk bool
}

func (i info) speak() string {
	fmt.Println("I am ", i.fname, i.age)
}

func main() {

	i1 := info{
		person: person{"Rachel",
			"Green",
		},
		age: 34,
		ltk: true,
	}

	i2 := person{
		person: person{"Monica", "Geller"},
	}

	p1 :=
		i1.speak()
	fmt.Println(p1)

}
