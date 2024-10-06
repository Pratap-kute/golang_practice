package main

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
	var sums []int

	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
