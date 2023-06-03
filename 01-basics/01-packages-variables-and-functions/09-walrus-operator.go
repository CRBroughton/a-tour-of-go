package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// type inferred
	// := is only availale inside functions
	k := 3

	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
