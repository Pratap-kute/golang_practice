package main

import "fmt"

func main() {
	m := map[string]int{
		"james":           32,
		"Miss moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["james"])

	v, ok := m["barn"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss moneypenny"]; ok {
		fmt.Println("This is the if Print", v)
	}

}
