package main

import (
	"time"
)

func talk(ch chan string) {
	for {
		ch <- "talky talk talk"
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan string)
	go talk(ch)

	for {
		select {
		case text := <- ch:
			println(text)
		}
	}
}
