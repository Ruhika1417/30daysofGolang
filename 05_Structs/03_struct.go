package main

import "fmt"

func main() {

	type sex struct {
		male   string
		female string
	}

	type info struct {
		name   string
		gender sex
		age    int
	}

	i1 := info{
		name:   "Ted Mosby",
		gender: sex{male: "male"}, //not assigning female will create a "" null value space, after male
		age:    44,
	}
	fmt.Println(i1, info.name)

}
