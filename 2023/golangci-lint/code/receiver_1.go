package main

type Tobi struct {}

func (t Tobi) DrinkCoffee() {
	println("gluck..gluck..gluck")
}

func main(){
	tobi := new(Tobi)
	tobi.DrinkCoffee()
}