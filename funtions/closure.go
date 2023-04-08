package main

import "fmt"

func main() {
	fmt.Println(incrementor()())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
