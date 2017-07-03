package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var result map[string]int = make(map[string]int)
	splitted := strings.Fields(s)
	for i := 0; i < len(splitted); i++ {
		word := splitted[i]
		val, present := result[word]
		if present {
			result[word] = val + 1
		} else {
			result[word] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
