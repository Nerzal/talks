package main 

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go increment(chan1)
	go increment(chan2)

	for {
		select {
		case i := <-chan1:
			println("channel 1:", i)
		case i := <-chan2:
			println("channel 2:",i)
		}
	}
}

func increment(channel chan int) {
	for i:=0;;i++{
		channel <- i
	}
}

