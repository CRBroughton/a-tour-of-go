package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// iterates of the pow slice
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
}
