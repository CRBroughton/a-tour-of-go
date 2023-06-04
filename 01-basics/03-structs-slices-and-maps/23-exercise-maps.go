package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// words is a slice of strings
	words := strings.Fields(s)
	// create a map with a key of the word
	// a value of the number of times
	// it appears in the initial string
	wordcount := make(map[string]int)

	// add up how many times that word appears
	for _, value := range words {
		wordcount[value]++
	}
	return wordcount
}

func main() {
	wc.Test(WordCount)
}
