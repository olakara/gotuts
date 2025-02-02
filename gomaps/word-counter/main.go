package main

import "strings"

func split(sentence string) []string {
	return strings.Fields(strings.ToLower(sentence))
}

func main() {
	counts := map[string]int{}
	sentence := "The brown fox jumps over the lazy dog"
	words := split(sentence)
	for _, word := range words {
		counts[word]++
	}
	for word, count := range counts {
		println(word, count)
	}
}
