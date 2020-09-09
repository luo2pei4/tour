package main

import "fmt"

// Sqrt is a test of Newton's method
func Sqrt(x float64) float64 {

	var z float64 = 1.0

	for v := 0; v < 10; v++ {

		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {

	fmt.Println(Sqrt(2))
}
