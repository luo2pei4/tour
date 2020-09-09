package main

import "fmt"

var pow = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

func main() {

	for idx, val := range pow {

		fmt.Printf("2**%d = %d\n", idx, val)
	}
}
