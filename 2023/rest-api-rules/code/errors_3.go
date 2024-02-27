package main

import (
	"fmt"
	"errors"
)

func main(){
	innerError := errors.New("inner error")

	err := fmt.Errorf("outer error: %w", innerError)
	println(err.Error())
}