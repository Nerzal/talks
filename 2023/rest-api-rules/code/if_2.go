package main

func main(){
	if isCondition(){
		println("Condition is met")
		return
	} 

	println("Condition is not met")
}

func isCondition() bool {
	return true
}