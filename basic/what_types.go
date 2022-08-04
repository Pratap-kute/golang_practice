package main

import (
	"fmt"
)

var a int

type jai int

var b jai

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//type conversion
	//it's not casting ****
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
}
