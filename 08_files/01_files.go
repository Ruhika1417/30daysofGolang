package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// creating a file using the os package.
	file, err := os.Create("./demofile.txt")
	checkNilerror(err)

	// adding the input
	content := "I'm the first line inside demofile.txt"
	len, err := io.WriteString(file, content)
	checkNilerror(err)
	fmt.Println("length is", len) //bytes length

	//reading  the data
	data, err := ioutil.ReadFile("./demofile.txt")
	checkNilerror(err)
	//fmt.Println(data)
	fmt.Println(string(data))

}

func checkNilerror(err error) {
	if err != nil {
		panic(err)
	}

}
