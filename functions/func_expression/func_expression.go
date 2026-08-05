package main

import "fmt"

func main() {
	fmt.Println("Hello from main")

	f := func() {
		fmt.Println("Hello from func expresion")
	}
	f()

	g := func(x int) {
		fmt.Println("valve form g:", x)
	}

	g(42)
}
