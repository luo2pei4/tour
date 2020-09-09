package main

import (
	"fmt"
	"math"
)

var counter int = 0

func sqrt(x float64) string {

	fmt.Println("x is ", x)

	if x < 0 {

		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main() {

	fmt.Println(sqrt(-4))
}
