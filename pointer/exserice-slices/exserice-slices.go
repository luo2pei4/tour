package main

import "fmt"

// Pic is a test
func Pic(dx, dy int) [][]uint8 {

	ss := make([][]uint8, dy)

	for i := 0; i < dy; i++ {

		s := make([]uint8, dx)

		for j := 0; j < dx; j++ {

			s[j] = uint8((i + j) / 2)
		}

		ss[i] = s
	}

	return ss
}

func main() {

	fmt.Println("Pic resutl: ", Pic(10, 10))
}
