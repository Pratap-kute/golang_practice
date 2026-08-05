/*
Write a program that prints a number in decimal, binary,and hex
*/
package main

import "fmt"

func main() {
	number := 2
	fmt.Printf("Number\tBinary\tHexadecimal\n")
	fmt.Printf("%d\t%b\t%#x", number, number, number)
}
