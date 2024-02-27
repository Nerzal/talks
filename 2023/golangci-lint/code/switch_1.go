package main

func main(){
	value := 3

	switch value {
	case 1:
		println("hit first case")
	case 2:
		println("hit second case")
	case 3:
		println("hit third case")
	default:
		println("your value is in another castle")
	}
}