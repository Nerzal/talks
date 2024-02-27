package main

import "fmt"

func main(){
	numbers := []int{1,2,3,4,5,6,7,8}
	
	numbers2 := make([]int, 0, 10)
	fmt.Println("Capacity: ", cap(numbers2), " Length: ", len(numbers2))

	numbers2 = append(numbers2, numbers...)
	fmt.Println("Capacity: ", cap(numbers2), " Length: ", len(numbers2))

	numbers2 = append(numbers2, 9)
	fmt.Println("Capacity: ", cap(numbers2), " Length: ", len(numbers2))

	numbers2 = append(numbers2, 2,3)
	fmt.Println("Capacity: ", cap(numbers2), " Length: ", len(numbers2))

}
