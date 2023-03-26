package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Hello from anonymous func")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}

func foo() {
	fmt.Println("Hello from foo")
}
