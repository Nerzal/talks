package main

func First[T any](xs []T) T {
	if len(xs) > 0 {
		return xs[0]
	}

	var zero T
	return zero
}

func main(){	
	println(First([]int{2,3,1}))
}