package main

import (
	"fmt"
	"math"
)

// MyFloat is a customize type
type MyFloat float64

// Vertex is a struct
type Vertex struct {
	X, Y float64
}

// Adder is an interface
type Adder interface {
	Abs() float64
	Test()
}

// Abs is a method of type MyFloat
func (f MyFloat) Abs() float64 {

	if f < 0 {

		return float64(-f)
	}

	return float64(f)
}

// Test is a method of type MyFloat
func (f MyFloat) Test() {

	fmt.Println("MyFloat...implement interface Adder's Test method.")
}

// Abs is a method of type Vertex
func (v *Vertex) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Test is a method of type Vertex
func (v Vertex) Test() {

	fmt.Println("Vertex...implement interface Adder's Test method.")
}

func main() {

	var i Adder

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	i = f
	fmt.Println(i.Abs())
	i.Test()

	i = &v
	fmt.Println(i.Abs())
	i.Test()
}
