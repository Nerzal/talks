package main

import (
	"errors"
	"fmt"
)

type SomeTypeError struct {}

func (s SomeTypeError) Error() string {
	return "yep, shit hit the fan!"
}

func main(){
	innerError := &SomeTypeError{}
	err := fmt.Errorf("Woopsie: %w", innerError)
	
	switch  errors.Unwrap(err).(type) {
		case *SomeTypeError:
			println("My error is of type: SomeTypeError")
		default:
			println("My error is in another castle!")
	}
}