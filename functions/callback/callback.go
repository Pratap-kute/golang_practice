package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// variadic function parameter
	s := sum(ii...)
	fmt.Println("sum of all number", s)

	s2 := even(sum, ii...)
	fmt.Println("sum of even number", s2)

}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//calculate sum of even number it takes a function as a parameter and variadic parameter
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

//calculate sum of odd number it takes a function as a parameter and variadic parameter
func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
