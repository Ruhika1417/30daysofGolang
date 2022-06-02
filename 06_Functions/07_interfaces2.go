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

func (p person) speak() {
	fmt.Println("I am", p.fname, p.lname)
}

func (i info) speak() {
	fmt.Println("I am", i.fname, i.lname, i.ltk)
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Person", h.(person).fname)
	case info:
		fmt.Println("Info", h.(info).fname)

	}

	fmt.Println("I'm passed in Bar", h)

}

func main() {

	i1 := info{
		person: person{"Rachel",
			"Green",
		},
		age: 34,
		ltk: true,
	}

	p1 := person{
		"Ross",
		"Geller",
	}

	fmt.Println(i1)
	i1.speak()
	p1.speak()

	bar(p1)
	bar(i1)
	//bar can take in both person and info. coz both are also of type  human interface
	//polymorphism
}
