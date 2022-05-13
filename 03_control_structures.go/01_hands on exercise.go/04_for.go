//Create a for loop using this syntax
//● for condition { }
//Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	x := 1
	age:=20
	for x <= age{

		fmt.Println(x)
		x++

	}
}


/*previous code

func main(){
	for {
		age:=20
		if x:=1;x<=age;x++{
			fmt.Println(x)
			x++
		}
	}

}