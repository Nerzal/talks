package main

import ("fmt")

func main(){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	numbers := []int{1,2,3,4}

	println(numbers[6])
}