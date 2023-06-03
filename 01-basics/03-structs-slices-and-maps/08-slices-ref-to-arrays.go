package main

import "fmt"

// A slice does not store any data,
// it just describes a section of an underlying array.

// Changing the elements of a slice modifies
// the corresponding elements of its underlying array.

// Other slices that share the same underlying array
// will see those changes.

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	names[0] = "testasd"

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// modifiying the slice
	// modifies the underlying array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
