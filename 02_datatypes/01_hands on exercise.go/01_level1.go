/* Using the following operators, write expressions and assign their values to variables:
g. ==
h. <=
i. >=
j. !=
k. <
l. >  */

package main

import "fmt"

func main() {
	a := 45
	b := 45

	g := (a == b)
	h := a <= b
	i := a >= b
	j := a != b
	k := a < b
	l := a > b
	fmt.Println(g, h, i, j, k, l)

}
