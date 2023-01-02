/*
Create and use an anonymous struct.
*/

package main

import "fmt"

func main() {
	ano := struct {
		first    string
		friends  map[string]int
		favDriks []string
	}{
		first: "james",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDriks: []string{
			"martini",
			"water",
		},
	}
	fmt.Println(ano)
}
