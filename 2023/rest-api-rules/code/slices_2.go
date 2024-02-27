package main

import "fmt"

func main(){
	numbers := []int{1,2,3}
	numbers2 := numbers
	numbers[0] = 123

	fmt.Println(numbers2)
}