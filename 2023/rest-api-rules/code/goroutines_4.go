package main

import (
	"time"
	"sync"
)

func longRunningTask(workerNumber int){
	println("Worker:", workerNumber, "started")
	time.Sleep(time.Second)
	println("Worker:", workerNumber, "done")
}

func main(){
	var wg sync.WaitGroup
	for i:= 1; i<= 5; i++ {
		wg.Add(1)

		i := i
		go func(){
			defer wg.Done()
			longRunningTask(i)
		}()
	}
	wg.Wait()
}