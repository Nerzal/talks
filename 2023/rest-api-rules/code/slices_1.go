package main

import "fmt"

func main(){
	numbers := []int{1,2,3}
	fmt.Println(numbers)
	fmt.Println("Capacity: ", cap(numbers), " Length: ", len(numbers))
}