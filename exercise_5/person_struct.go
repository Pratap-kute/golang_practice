/*
Create your own type “person” which will have an underlying type of “struct” so that it can storethe following data:
●first name
●last name
●favorite ice cream flavors
Create two VALUES of TYPE person.
Print out the values, ranging over the elements in the slicewhich stores the favorite flavors.
*/

package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "amol",
		lastName:  "raj",
		iceCreamFlavors: []string{
			"vanilla",
			"strawberry",
		},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, j := range p1.iceCreamFlavors {
		fmt.Println(i, j)
	}

}
