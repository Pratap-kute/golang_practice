/*
Build and use an anonymous func
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello from anonymous func")
	}()

	fmt.Println("Hello from main func")
}
