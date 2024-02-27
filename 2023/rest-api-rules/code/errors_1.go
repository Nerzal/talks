package main

import "errors"

func someFunc() error {
	return errors.New("Wooopsie, sorry something went wrong :(")
}

func main(){
	err := someFunc()
	if err != nil {
		println(err.Error())
		return
	}

	println("Hooray, no errors, lets ROCK!")
}