/*
Using iota, create 4 constants for the NEXT4 years.Print the constant values.
*/
package main

import "fmt"

const (
	a = 2022 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
