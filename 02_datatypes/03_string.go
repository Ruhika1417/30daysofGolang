package main

import "fmt"

func main() {
	a := "ruhika"
	fmt.Printf("%T %v\n", a, a)

	abyte := []byte(a) //string is a sequence of bytes.
	fmt.Println(abyte)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%#U", a[i]) //unicode utf8 char
	}
	fmt.Println(" ")

	fmt.Println("Range and value loop ")

	for i, v := range a {
		fmt.Println(i, v)

	}
	fmt.Printf("%q\n", a)  //string w quotes.
	fmt.Printf("% X\n", a) //  base16, 2 char per byte. uppercase

	fmt.Printf("% x\n", a) //way 2 byte loop; base16, 2 char per byte.

}
