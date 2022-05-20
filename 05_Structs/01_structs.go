//1.struct type(blueprint for struct) 2.struct. (actual struct)

package main

import "fmt"

func main() {

	type info struct {
		name    string
		age, id int
	}
	//creating struct or values of that type
	//compositer literal: type{}

	i1 := info{
		name: "Ruhika",
		age:  20,
		id:   2222, //order does not matter
	}

	i2 := info{
		age:  26,
		id:   2929,
		name: "Shailesh",
	}

	//way 2 of creating a struct
	//this is by creating an empty struct i3 and assigning values to it's fields individually
	var i3 info
	i3.name = "Ruhani"
	i3.age = 29
	i3.id = 1000

   //way3:without declaring the field names
 
   i4 :=info{
   
   

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)

	fmt.Println(i1.name)
	fmt.Println(i2.name)

}
