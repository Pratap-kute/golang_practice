package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
			age:   32,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.first)
}
