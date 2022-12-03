package main

import (
	"sort"
	"strings"
)

func main() {
	print(frequencySort("bbaaaaccccccc"))
}

func frequencySort(s string) string {
	var result string
	freq := make(map[rune]int)
	for _, i := range s {
		freq[i] = freq[i] + 1
	}
	var keys []rune
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})
	for _, c := range keys {
		result += strings.Repeat(string(c), freq[c])
	}
	return result
}
