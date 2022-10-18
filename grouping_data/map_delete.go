package main

import "fmt"

func main() {
	m := map[string]int{
		"james":           32,
		"miss moneypenny": 27,
	}
	fmt.Println(m)
	delete(m, "james")
	fmt.Println(m)
}
