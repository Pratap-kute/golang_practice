package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "amol",
		last:  "Bond",
	}
	p2 := person{
		first: "jack",
		last:  "momo",
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.first)
}
