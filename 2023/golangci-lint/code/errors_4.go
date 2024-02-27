package main

import (
	"fmt"
	"errors"
)

func calcPi() error {
	return errors.New("calculate pi")
}

func main(){
	calcPiErr := calcPi()

	err := fmt.Errorf("processing request: %w", calcPiErr)
	println(err.Error())
}