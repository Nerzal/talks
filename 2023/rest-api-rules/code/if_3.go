package main

func main(){
	if condition := isCondition(); condition {
		println("Condition is met with value:", condition)
		return
	} 

	println("Condition is not met")
}

func isCondition() bool {
	return true
}