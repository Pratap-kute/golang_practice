/*
Hands-on exercise #3
Create a for loop using this syntax
	● for condition { }
Have it print out the years you have been alive.
*/

package main

import "fmt"

func main() {
	birth_year := 1993
	for birth_year <= 2022 {
		fmt.Println(birth_year)
		birth_year++
	}
}
