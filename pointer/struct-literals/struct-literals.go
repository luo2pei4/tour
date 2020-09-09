package main

import "fmt"

// Vertex is a struct
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	v4 = &Vertex{1, 2}
)

func main() {

	fmt.Println(v1, v2, v3, v4)
}
