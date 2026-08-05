package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 42
	b := 43
	//Expression evaluation means that the value of the expression after comparison
	fmt.Println(a == b)
}
