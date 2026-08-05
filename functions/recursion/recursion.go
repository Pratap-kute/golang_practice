/*
factorial of a given integer using a recursive function.

The main function calls the factorial function and passes the argument 5 to it.
The factorial function then takes this argument and calculates its factorial.

The factorial function first checks if the input is equal to 0.
If it is, then it returns 1, since the factorial of 0 is defined as 1.
If the input is not 0, then it calculates the factorial recursively by multiplying the input by the factorial of the input minus 1.
The recursion continues until the input is reduced to 0.

So in this case, the factorial function is called with input 5.
Since 5 is not equal to 0, the function calculates the factorial by multiplying 5 with the factorial of 4,
which is calculated by calling the factorial function again with input 4.
This continues until the input is reduced to 0, at which point the function returns 1, and the recursion stops.

The final result is then printed by the fmt.Println function in the main function, which outputs the value 120, which is the factorial of 5.

*/

package main

import "fmt"

func main() {
	fmt.Println(factorial(7))
	fmt.Println(factorial_for(7))
	fmt.Println(factorial_for(7))

}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func factorial_for(x int) int {
	var result int = 1
	for i := 1; i <= x; i++ {
		result *= i
	}
	return result
}

func loofact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
