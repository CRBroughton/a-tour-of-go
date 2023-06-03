package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// can declare the variable in-line on the if statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
