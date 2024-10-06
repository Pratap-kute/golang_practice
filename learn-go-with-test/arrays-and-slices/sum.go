package main

import "fmt"

func Sum(number []int) int {
	sum := 0
	// we are using range returns two values - the index and the value.
	// We are choosing to ignore the index value by using _ blank identifier.
	for _, number := range number {
		sum += number
	}
	return sum
}

// "...[]int" is used to accept the any lenght of slice
func SumAll(numberToSum ...[]int) []int {
	lengthOfNumber := len(numberToSum)
	capasityofNumber := cap(numberToSum)
	fmt.Printf("capacity of slice %d \n", capasityofNumber)
	// make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through.
	sums := make([]int, lengthOfNumber)

	for i, numbers := range numberToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
