package main
import (
	"context"
	"time"
)
func talk(ctx context.Context, ch chan string) {
	for {
		select {
		case ch <- "talky talk talk":
			time.Sleep(time.Second)
		case <-ctx.Done():
			println("Er ist tod, jim!")
			return
		}
	}
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second *5)
	ch := make(chan string)
	go talk(ctx, ch)

	for {	
		println(<- ch)
	}
}
