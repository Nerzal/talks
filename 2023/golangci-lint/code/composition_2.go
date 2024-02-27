package main

type Vector2 struct { 
}

func (v Vector2) Print() { println("X: 2 Y: 3") }

type Vector3 struct {
	Vector2
}

func main() {
	vector3 := Vector3{}
	vector3.Print()
}