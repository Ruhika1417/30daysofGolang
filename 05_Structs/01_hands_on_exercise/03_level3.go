// Hands-on exercise #3
// ● Create a new type: vehicle.
// ○ The underlying type is a struct.
// ○ The fields:
// ■ doors
// ■ color
// ● Create two new types: truck & sedan.
// ○ The underlying type of each of these new types is a struct.
// ○ Embed the “vehicle” type in both truck & sedan.
// ○ Give truck the field “fourWheel” which will be set to bool.
// ○ Give sedan the field “luxury” which will be set to bool. solution
// ● Using the vehicle, truck, and sedan structs:
// ○ using a composite literal, create a value of type truck and assign values to the
// fields;
// ○ using a composite literal, create a value of type sedan and assign values to the
// fields.
// ● Print out each of these values.
// ● Print out a single field from each of these values.

package main

import "fmt"

func main() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourwheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}
	//mistake: t1:=struct{}

	// t1 := truck{
	// 	doors:     4,
	// 	color:     "black",
	// 	fourwheel: true,
	// }
	t1 := truck{
		vehicle:   vehicle{doors: 4, color: "Black"},
		fourwheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{doors: 4, color: "Black"},
		luxury:  true,
	}

	fmt.Println(t1, s1)

}
