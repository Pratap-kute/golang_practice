package main

import "fmt"

// when you shift binary digits to the left or right. study more about it
func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}
