package main

import "fmt"

func main(){
	numbers := []int{1,2,3}
	numbers2 := numbers

	foo(numbers)

	fmt.Println(numbers2)
}

func foo(numbers []int) {
	numbers[1] = 42
}