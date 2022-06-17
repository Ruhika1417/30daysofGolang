package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Rachel",
		Last:  "Green",
		Age:   27,
	}
	p2 := Person{
		First: "Ross",
		Last:  "Geller",
		Age:   27,
	}
	fmt.Println(p1)

	people := []Person{p1, p2}
	fmt.Println(people)

	byteb, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteb))
}
