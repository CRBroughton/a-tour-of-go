package main

import "fmt"

// init the variables
var i, j int = 1, 2

func main() {
	// init the variables, inferred
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
