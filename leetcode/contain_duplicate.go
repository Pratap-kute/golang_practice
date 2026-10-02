package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, val := range nums {
		if seen[val] {
			return true
		}
		seen[val] = true
	}

	fmt.Println(seen)

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}

	duplicate := containsDuplicate(nums)
	fmt.Println(duplicate)
}
