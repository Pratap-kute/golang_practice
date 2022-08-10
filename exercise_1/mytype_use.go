/*
1.at the package level scope, using the “var” keyword,
create a VARIABLE with the IDENTIFIER “y”.
The variable should be of the UNDERLYING TYPE of your customTYPE “x”
	a.eg:
		type hotdog int
		var x hotdog
		var y int
2.in func main
	a.this should already be done
		i.print out the value of the variable “x”
		ii.print out the type of the variable “x”
		iii.assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
*/
package main

import "fmt"

type myType int
var x myType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y=int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
