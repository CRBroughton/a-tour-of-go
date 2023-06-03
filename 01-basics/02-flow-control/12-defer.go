package main

import "fmt"

func main() {
	// this is 'deferred' and runs after
	// the main funtion returns
	defer fmt.Println("world")

	fmt.Println("hello")
}
