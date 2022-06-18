package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type person struct {
		first string `json:"first"`
		last  string `json:"last"`
	}

	j1 := `[{"first":"rachel", "last": "green"},{"first":"ross", "last": "geller"}]`
	bs := []byte(j1)

	//people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		panic(err)
	}

	fmt.Println(people)

}
