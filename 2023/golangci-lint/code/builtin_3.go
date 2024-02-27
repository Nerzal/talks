package main

import ("fmt")

func main() {
	myMapping := map[string]int {
		"pi": 3,
		"e": 3,
		"sqrtMinus2": 1,
	}

	fmt.Println(myMapping)

	delete(myMapping, "pi")

	fmt.Println(myMapping)
}