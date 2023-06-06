package main

import "fmt"

type I interface {
	M()
}

func main() {
	// methods called on this
	// interface will cause a
	// run-time error
	// as it is nil
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
