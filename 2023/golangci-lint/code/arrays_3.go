package main

import "fmt"

func main(){
	numbers := [4]int{}
	fmt.Println("Numbers: ", numbers)

	strings := [3]string{}
	fmt.Println("Strings: ", strings)

	others := [5]interface{}{}
	fmt.Println("Others: ", others)
}
