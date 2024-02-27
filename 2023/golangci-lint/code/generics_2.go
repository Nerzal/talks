package main

type LinkedList[T any] struct {
	value T
	next *LinkedList[T]
}

func (list *LinkedList[T]) Get() T {
	return list.value
}

func main(){
	list := &LinkedList[int] {
		value: 4321,
	}

	println(list.Get())
}