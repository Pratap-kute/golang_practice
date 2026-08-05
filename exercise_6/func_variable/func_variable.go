/*
Assign a func to a variable, then call that func
*/

package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello from anonymous func")
	}
	f()

	fmt.Println("Hello from main func")
}
