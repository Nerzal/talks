package main

func main(){
	value := 1

	switch value {
	case 1:
		println("hit first case")
		fallthrough
	case 2:
		println("hit second case")
	case 3:
		println("hit third case")
	}
}