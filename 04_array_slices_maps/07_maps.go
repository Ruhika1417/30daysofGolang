package main

import "fmt"

var m = map[string]int{
	"A": 65,
	"B": 66,
	"C": 67,
	"Z": 90,
}

func main() {

	//can access elements by keys
	fmt.Println(m)
	fmt.Println(m["A"])

	fmt.Println(m["D"]) //returns a null value

	v, ok := m["D"] //to check if it is actually present or not with comma idiom.
	fmt.Println("v is ", v, " ,", "ok is", ok)

	if v, ok := m["D"]; ok {
		fmt.Println("This will not be printed coz the statement is false", v)
	}

}
