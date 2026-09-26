package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var target_index []int
	for index, val := range nums {
		// find_this_number := target - val

		for i, v := range nums {
			if v+val == target && index != i {
				target_index = append(target_index, index, i)
				break
			}
		}
		if len(target_index) == 2 {
			break
		}
	}

	return target_index
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	indexoftwosum := twoSum(nums, target)
	fmt.Println(indexoftwosum)
}
