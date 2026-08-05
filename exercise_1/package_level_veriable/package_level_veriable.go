/*
1.Use var toDECLAREthreeVARIABLES.The variablesshould have package levelscope.
Do not assignVALUESto the variables.Use the following IDENTIFIERSfor the variables and make sure
the variables are of the following TYPE (meaning they canstore VALUES of thatTYPE).
	a.identifier “x” type int
	b.identifier “y” type string
	c.identifier “z” type bool
2.in func main
	a.print out the values for each identifier
	b.The compiler assigned values to the variables.
	What are these values called?
	ANS : zero values
*/
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
