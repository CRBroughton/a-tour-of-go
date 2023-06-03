package main

import "fmt"

func main() {
	// without an init value
	// the zero value will be used ("", 0, false)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %b %q\n", i, f, b, s)
}
