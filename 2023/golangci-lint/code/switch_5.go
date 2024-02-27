package main

type someInterface interface {}
type someType struct {}
type someOtherType struct {}

func main(){
	var value someInterface
	value = someType{}

	switch value.type() {
	case someType:
		println("value is someType")
	case someOtherType:
		println("value is someOtherType")
	}
}