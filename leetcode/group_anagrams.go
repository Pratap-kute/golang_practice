package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	var groupAnagrams [][]string
	anagram := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(str)
		// fmt.Println(chars)
		slices.Sort(chars)
		anagram[string(chars)] = append(anagram[string(chars)], str)

	}

	for _, group := range anagram {
		groupAnagrams = append(groupAnagrams, group)
	}

	// fmt.Println(anagram)

	return groupAnagrams
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
