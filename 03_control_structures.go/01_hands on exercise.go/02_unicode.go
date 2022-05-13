/* Print every rune code point of the uppercase alphabet three times. Your output should look like
this:
65
U+0041 'A'
U+0041 'A'
U+0041 'A'
*/
package main

import "fmt"

func main() {

	for i := 65; i < 92; i++ {
		fmt.Println(i)
		fmt.Printf("\t%#U\n \t%#U\n \t%#U\n", i, i, i)
	}

}
