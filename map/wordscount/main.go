package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words_count := make(map[string]int)
	string_array := strings.Fields(s)
	for i := 0; i < len(string_array); i++ {
		words_count[string_array[i]] += 1
	}
	return words_count
}

func main() {
	wc.Test(WordCount)
}
