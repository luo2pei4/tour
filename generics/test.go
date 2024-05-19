package main

import "fmt"

type CommonType[T int | string | float32] []T

func main() {
	fmt.Println("hello world")
}
