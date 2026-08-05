/*
Create a slice of a slice of string ([][]string).
Store the following data in the multi-dimensionalslice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny","Helloooooo, James."
Range over the records, then range over the data in each record.
*/

package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	sample := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	fmt.Printf("Number of rows in array: %d\n", len(sample))
	fmt.Printf("Number of columns in array: %d\n", len(sample[0]))
	fmt.Printf("Total number of elements in array: %d\n", len(sample)*len(sample[0]))

	multi := [][]string{x, y}
	fmt.Println(multi)

	for i, b := range multi {
		fmt.Println("recod: ", i)
		for j, c := range b {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, c)
		}
	}
}
