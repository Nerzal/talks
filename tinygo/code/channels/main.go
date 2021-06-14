package main

var myChannel chan int

func main() {
	myChannel = make(chan int)

	for {
		println(<-myChannel)
	}
}

func asyncFunction(myChannel chan int) {
	i := 0
	for {
		myChannel <- i
		i++
	}
}
