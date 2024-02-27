package main

type Vector2 struct { 
	X int
	Y int 
}

type Vector3 struct {
	Vector2
	Z int
}

func main() {
	vector3 := Vector3{
		Vector2: Vector2{X: 1, Y: 2},
		Z: 4,
	}

	println("X:", vector3.X, "Y:", vector3.Y, "Z:", vector3.Z)
}