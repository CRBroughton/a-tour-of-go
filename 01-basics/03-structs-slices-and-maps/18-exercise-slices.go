package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	array := make([][]uint8, dy)

	for i := range array {
		array[i] = make([]uint8, dx)
	}
	return array

}

func main() {
	pic.Show(Pic)
}
