package main

type Drinker interface { Drink() }

type Tobi struct {}

func (t Tobi) Drink() { println("gluck gluck...bööörp") }

type Artur struct {}

func (a Artur) Drink() { println("gluck gluck") }

func main(){
	var mensch Drinker

	mensch = Tobi{}

	mensch.Drink()
}
