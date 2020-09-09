package main

import "fmt"

// ErrNegativeSqrt is a float64 type
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt is a function
func Sqrt(x float64) (float64, error) {

	if x < 0 {

		return float64(x), ErrNegativeSqrt(x)
	}

	var z float64 = 1.0

	for v := 0; v < 10; v++ {

		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-3))
}
