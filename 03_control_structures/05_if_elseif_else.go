package main

import "fmt"

func main() {

	a := 131

	if a%5 == 0 {
		fmt.Println(a, "is divisible by 5")

	} else if a%6 == 0 {
		fmt.Println(a, "is divisble by 6")

	} else if a%7 == 0 {
		fmt.Println(a, "is divisble by 6")
	} else {
		fmt.Println(a, "is divisible by any other number or is a prime no")
	}

	//priniting out even numbers loop + condition
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)

		}
	}

}
