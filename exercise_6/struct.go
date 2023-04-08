/*
● Create a user defined struct with
	○ the identifier “person”
	○ the fields:
		■ first
		■ last
		■ age
● Attach a method to type person with
	○ the identifier “speak”
	○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/
package main

import "fmt"

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p1.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}
