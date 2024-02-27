package main

import "fmt"

func main(){
	numbers := []int{1,2,3,4,5,6,7,8}
	
	fmt.Println("Die ersten 3 Elemente sind: ", numbers[:3])

	fmt.Println("Die letzten 4 Elemente sind:", numbers[4:])

	fmt.Println("Slicy Slice - Element zwischen 2 und 4:", numbers[2:4])
}
