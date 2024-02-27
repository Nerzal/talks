package main

import (
	"fmt"
	_ "embed"
	"encoding/json"
)

//go:embed assets/coolStuff.json
var data []byte

func main() {
	var myData map[string]string

	err := json.Unmarshal(data, &myData)
	if err != nil {
		panic(err)
	}

	fmt.Println(myData)
}