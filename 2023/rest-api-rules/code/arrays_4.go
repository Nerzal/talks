package main

import "fmt"

func main(){
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1

	arr1[0] = 1337

	fmt.Println(arr2)
}
