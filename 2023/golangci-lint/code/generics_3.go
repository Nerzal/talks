package main

func First(xs []int) int {
	if len(xs) > 0 {
		return xs[0]
	}

	
	return 0
}

func main(){	
	println(First([]int{2,3,1}))
}