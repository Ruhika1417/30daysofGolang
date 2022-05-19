package main

import "fmt"

func main() {

	m := map[string]int{
		"Robert": 678,
		"Ken":    675,
		"Rob":    890,
	}
	fmt.Println(m)
	fmt.Println(" ")

	//1.adding elements to a map

	m["Todd"] = 56
	fmt.Print(m, "\n\n")

	//2.deletion
	delete(m, "Tod") //Tod not present
	delete(m, "Todd")
	fmt.Print(m, "\n\n")

	//3.ranging over a map

	for i, v := range m {
		fmt.Println(i, v)
	}

}
