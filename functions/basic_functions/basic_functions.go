package main

import (
	"fmt"
)

func main() {
	foo()
	slc := []int{11, 22, 33, 44, 55, 66}
	foo1(23, 54, 34634, 643, 12)
	// unpacking the slice ***very important to understant the con
	foo1(slc...)
}

func foo() {
	fmt.Println("Hello,playground")
}

// packing the slice using ...
func foo1(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

// unpacking the slice
