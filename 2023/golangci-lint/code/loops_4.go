package main

func main(){
	names := []string{"Alice", "Bob", "Chad", "Dave", "Erin"}

	for index, value := range names {
		println("index:", index, "value:", value)
	}
}