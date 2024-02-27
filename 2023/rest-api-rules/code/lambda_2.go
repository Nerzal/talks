package main

import ("time")

func main() {
	coolFunc := func() {
		for {
			println("Mrs. cool rennt in einer anonymen function")
			time.Sleep(time.Second)
		}
	}

	go coolFunc()

	select {}
}