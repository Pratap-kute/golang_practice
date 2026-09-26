package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var target_index []int
	seen := make(map[int]int)
	for index, val := range nums {
		find_this_number := target - val
		ind, found := seen[find_this_number]

		if found {
			target_index = append(target_index, ind, index)
			break
		} else {
			seen[val] = index
		}

	}

	return target_index
}

func main() {
	nums := []int{3, 2, 4}
	target := 7

	indexoftwosum := twoSum(nums, target)
	fmt.Println(indexoftwosum)
}
