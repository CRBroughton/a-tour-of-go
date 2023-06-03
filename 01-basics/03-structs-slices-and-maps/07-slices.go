package main

import "fmt"

func main() {
	// a fixed sized array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// a dynamically sized slice
	// formed from two indices, low and high
	// excludes the highest in the range [3 5 7]
	var s []int = primes[1:4]
	fmt.Println(s)
}
