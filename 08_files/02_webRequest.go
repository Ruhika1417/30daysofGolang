package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://lco.gg/"

	response, err := http.Get(url)
	checkNilerror(err)

	//resonse will be in the form of header and body.
	fmt.Printf("%T\n", response)
	fmt.Println(response.StatusCode)

	data, err := ioutil.ReadAll(response.Body)
	checkNilerror(err)

	fmt.Println(string(data))

	defer response.Body.Close() //important coz a request never closes by itself

}

func checkNilerror(err error) {
	if err != nil {
		panic(err)
	}
}
