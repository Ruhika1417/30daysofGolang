package main

import "fmt"

func main() {
	//1
	var x [2]int
	fmt.Println(x)
	x[0] = 10
	x[1] = 11
	fmt.Println(x)
	fmt.Printf("%T %v\n", x, x) //to check type

	//2
	array1 := [...]int{2, 3}
	fmt.Printf("this is array1 %T %v\n", array1, array1)

	//3

	array2 := [2][2]byte{
		{1, 2}, {3, 4}}
	fmt.Printf("this is array2 %T %v\n", array2, array2)

	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			fmt.Printf("array2[%d[%d] =%d\n", i, j, array2[i][j])
		}
	}

}
