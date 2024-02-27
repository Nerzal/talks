package main

func main() {
	number := 3 // Also exakt die Eulerische Zahl

	doSomething(&number)

	println(number)
}

func doSomething(number *int) {
	*number += 1
}