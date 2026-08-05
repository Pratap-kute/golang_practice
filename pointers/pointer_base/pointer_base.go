package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)

	// & gives you the address and different from * which gives you the value stored at an address
	// CAN NOT ASSING A POINTER TO A NON POINTER MEANS C= &A IS NOT VALID
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	// * gives you the value stored at an address when you have the address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
}
