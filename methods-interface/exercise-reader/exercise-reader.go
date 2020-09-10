package main

import "fmt"

// MyReader is a struct
type MyReader struct{}

// Read, implements Reader interface
func (mr MyReader) Read(b []byte) (int, error) {

	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {

	fmt.Println(MyReader{})
}
