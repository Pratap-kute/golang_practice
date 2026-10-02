package main

import "fmt"

func mpas_study() {
	seen := make(map[int]int)

	seen[3] = 0
	seen[2] = 1

	index, exists := seen[7]
	fmt.Println(index, exists)

	index, exists = seen[2]
	fmt.Println(index, exists)
}

func main() {
	mpas_study()
}
