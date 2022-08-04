package main

import (
	"fmt"
)

var y string
var z int

func main() {
	//Declare a variable to be of a certain type
	//and then assign a value of that type to the variable
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	y = "Bond, James Bond"
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
