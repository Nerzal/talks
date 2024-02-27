package main

type Person struct { Age int }

func (p Person) DoAge() {
	p.Age += 1
}

func main(){
	tobi := Person { Age : 30}
	tobi.DoAge()

	println(tobi.Age)
}