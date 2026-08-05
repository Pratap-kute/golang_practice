/*
Print out the remainder (modulus) which is found for each number between 10 and 100 when itis divided by 4.
*/
package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		mod := i % 4
		fmt.Printf("%v reminder of 4 is %v\n", i, mod)
	}
}
