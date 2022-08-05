/*
Using the code from the previous exercise,
1.At the package level scope, assign the following values to the three variables
	a.for x assign 42
	b.for y assign “James Bond”
	c.for z assign true
2.in func main
	Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page26
	a.use fmt.Sprintf to print all of the VALUES to one single string.
		ASSIGN thereturned value of TYPE string using the short declaration operator to
		a VARIABLE with the IDENTIFIER “s”
	b.print out the value stored by variable “s”
*/
package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//%v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
