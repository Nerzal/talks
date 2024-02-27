package main

import ("time")

func main() {
	go func() {
		for {
			println("Mrs. cool rennt in einer anonymen function")
			time.Sleep(time.Second)
		}
	}()

	select {}
}