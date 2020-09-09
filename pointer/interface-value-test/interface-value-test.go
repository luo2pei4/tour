package main

import "fmt"

type it interface {
	NextValue(x int)
}

type S struct {
	A int
}

func (s S) NextValue() {
	s.A = s.A + 1
}

func main() {

	s := S{0}

	for i := 0; i < 10; i++ {

		s.NextValue()
		fmt.Println(s.A)
	}
}
