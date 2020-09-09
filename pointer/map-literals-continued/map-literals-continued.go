package main

import "fmt"

// Vertex is a struct type
type Vertex struct {
	Lat, Long float64
}

// 此处在给map赋值的时候省略了Vertex类型说明,可对比注释部分
var m = map[string]Vertex{

	// "Bell Labs" : Vertex{40.68433, -74.39967},
	// "Google" : Vertex{37.42202, -122.08408},

	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
