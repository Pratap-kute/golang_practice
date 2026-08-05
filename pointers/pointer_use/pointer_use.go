package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	// change the value at this address to 43 using the pointer
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
