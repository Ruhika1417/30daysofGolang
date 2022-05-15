package main

import "fmt"

func main() {

	n := 3
	switch n {

	case 1:
		fmt.Println("1")
		//fallthrough

	case 2:
		fmt.Println("2")

	case 3:
		fmt.Println("3")
		fallthrough //fallthrough's job is to print the next statement as well if the current one is true(let the next one be a default one , that too is executed)

	default:
		fmt.Println("no switch case was executed or fallthrough was used in the previous case. ")

	}

	//if no expression is put in switch it will take true by deafult and will search for true cases.
	// if you fall through and use fallthrough consecutively even if the conditions are false it will be printed.

	switch {
	case true:
		fmt.Println("this should not be printed")

	case 2 == 2:
		fmt.Println("*******condition satisfied, hence should  be printed*****")
		fallthrough

	case (3 == 2):
		fmt.Println("printed bc of fallthrough")
		fallthrough

	case (3 == 1):
		fmt.Println("printed bc of fallthrough")

	default:
		fmt.Println("a default statement!")

	}

	//eg 3 switch on value

	switch "K" {
	case "B":
		fmt.Println("not R")
	case "R", "K", "g", "M": //multiple values example
		fmt.Println(`One of the values "R", "K", "g", "M" `)
	}

}
