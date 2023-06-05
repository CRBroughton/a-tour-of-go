package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// This is a 'method' attacted
// to the Vertex struct;
// This method can only be called
// on this type of struct
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
