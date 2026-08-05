/*
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:
make([]T,length,capacity)
capacity- this mean that it created the sized(length) ahead of time
*/
package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42

}
