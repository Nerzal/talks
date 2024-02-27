package main

import ("fmt")

func main() {
	numbers1 := []int{1,1,2,3,5,8,13,21}
	numbers2 := make([]int, len(numbers1))	

	numberCopied := copy(numbers2, numbers1)
	println("copied", numberCopied, "elements")

	fmt.Println(numbers2)
}