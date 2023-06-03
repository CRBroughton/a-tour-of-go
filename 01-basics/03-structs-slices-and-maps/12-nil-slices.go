package main

import "fmt"

func main() {
	// a nill slice
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
