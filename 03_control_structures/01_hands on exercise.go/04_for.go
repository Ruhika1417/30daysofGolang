//Create a for loop using this syntax
//● for condition { }
//Have it print out the years you have been alive.

package main

import "fmt"

func main() {

	born_y := 2001
	current := 2022

	for {
		if born_y <= current {
			fmt.Println(born_y)
		}
		born_y++

	}

}

/*previous code.

func main(){
	for {
		age:=20
		if x:=1;x<=age;x++{
			fmt.Println(x)
			x++
		}
	}

}
*/
