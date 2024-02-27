package main

import "time"

func main(){
	for i := 0;; i++{
		println("yep yep yep yep yep yep yep yep yep")
		time.Sleep(time.Second)

		if i == 3 {
			break
		}
	}
}