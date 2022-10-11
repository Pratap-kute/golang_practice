// here we are using range to print index and value
package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 54}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])

	//using range to pront index and value
	for i, v := range x {
		fmt.Println(i, v)
	}
}
