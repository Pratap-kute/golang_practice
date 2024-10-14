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
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices
func SumAllTails(numberToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
