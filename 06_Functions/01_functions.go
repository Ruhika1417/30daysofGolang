package main

import "fmt"

func hello() {
	fmt.Println("hello world")
}

//we define our func w parameters(if any)
func func_name(name string) {
	fmt.Println("My name is", name)
}

//the scope of the identifier:name is till end of func only },so it's to have two para w same name in diff func
func woo(name string) string {
	return fmt.Sprint("Hello, ", name)
}

//4. returns multiple values

func name(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, " says hello")
	b := true
	return a, b
}

func main() {

	hello()
	//we call our func with arguements(if any)
	func_name("James")

	f2 := woo("Jacqueline")
	fmt.Println(f2)

	j, f := name("Jacqueline ", "Fernandez")
	fmt.Println(j)
	fmt.Println(f)

}
