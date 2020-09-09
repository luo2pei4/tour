package main

import "fmt"

func printSlice(s []int) {

	fmt.Printf("addr=%p len=%d cap=%d %v\n", s, len(s), cap(s), s)
}

func main() {

	// var arr = []int{1, 2, 3, 4, 5, 6}

	// fmt.Printf("arr first element addr: %p\n", arr)

	// var slc = arr[:]

	// fmt.Printf("addr of slice: %p\n", slc)\

	// slc = arr[1:]

	// fmt.Printf("addr of slice: %p\n", slc)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[1:5]
	printSlice(s)

	s = s[2:3]
	printSlice(s)

}
