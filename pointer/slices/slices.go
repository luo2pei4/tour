package main

import "fmt"

func main() {

	arr := [6]int{1, 2, 3, 4, 5, 6}
	var slice []int = arr[1:4]
	fmt.Println(slice)
}
