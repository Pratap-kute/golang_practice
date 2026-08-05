/*
 ● Using a COMPOSITE LITERAL:
    ● create a SLICE of TYPE int
    ● assign 10 VALUES
    ● Range over the slice and print the values out.
    ● Using format printing
        ○ print out the TYPE of the slice
*/
package main

import "fmt"

func main() {
	x := []int{5, 7, 8, 9, 11, 23, 89, 45, 50}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
