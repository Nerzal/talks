package main

import "fmt"

func main(){
	numbers := createArray(3)
	fmt.Println(numbers)
}

func createArray(size int) []int {
	numbers := [size]int{}

	return numbers
}