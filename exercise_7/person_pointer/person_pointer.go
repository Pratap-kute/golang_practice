/*
● create a person struct
● create a func called “changeMe” with a *person as a parameter
	○ change the value stored at the *person address
		■ important
● to dereference a struct, use (*value).field
	○ p1.first
	○ (*p1).first
● “As an exception, if the type of x is a named pointer type and (*x).
fis a valid selector expression denoting a field (but not a method),
x.f is shorthand for (*x).f.”
	○ https://golang.org/ref/spec#Selectors
● create a value of type person
	○ print out the value
● in func main
	○ call “changeMe”
● in func main
	○ print out the value
*/

package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	address    string
}

func changeMe(p *person) {
	p.first_name = "amol"
	p.last_name = "raj"
	p.address = "pune"
}

func main() {
	p1 := person{
		first_name: "james",
		last_name:  "bond",
		address:    "london",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
