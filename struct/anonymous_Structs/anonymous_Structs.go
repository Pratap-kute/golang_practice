package main

import "fmt"

func main() {
//    anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1.first)
}
