package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// stacked defers go last-in-first-out
		// https://go.dev/blog/defer-panic-and-recover
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
