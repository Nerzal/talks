package main

type SomeTypeError struct {}

func (s *SomeTypeError) Error() string {
	return "yep, shit hit the fan!"
}

func main(){
	err := &SomeTypeError{}
	if err != nil {
		println(err.Error())
	}
}