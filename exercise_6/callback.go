/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
	● pass a func into a func as an argument


The g function takes an integer slice as an argument and returns the sum of the first and last elements of the slice.
If the slice is empty, it returns 0. If the slice has only one element, it returns the value of the element.

The callback function takes a function that expects an integer slice and returns an integer and an integer slice as arguments.
It calls the provided function passing the integer slice as an argument and increments the resulting value by one.
It then returns the incremented value.

In the main function, a closure function g is defined and passed as an argument to
the callback function along with an integer slice {1, 2, 3, 4, 5}.
The callback function is called and returns the result of the g function plus one. The result is assigned to the variable x.
The x value is then printed to the console with the message "The total is".

In summary, the program calculates the sum of the first and last elements of an integer slice and increments
the result by one using a callback function. The final result is printed to the console.

*/

package main

import "fmt"

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := callback(g, []int{1, 2, 3, 4, 5})
	fmt.Println("The total is", x)

}

func callback(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
