package main

import "fmt"

func printSlice(s []int) {

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {

	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)

	e := []int{2, 3, 5, 7, 11}
	f := e[:2]
	printSlice(f)

	g := f[2:5]
	printSlice(g)
}
