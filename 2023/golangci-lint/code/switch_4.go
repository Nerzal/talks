package main

func main(){
	value := 3

	switch {
	case value == 1:
		println("hit first case")
	case value == 2:
		println("hit second case")
	case value == 3:
		println("hit third case")
	default:
		println("your value is in another castle")
	}
}