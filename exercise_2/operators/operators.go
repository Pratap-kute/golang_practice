/*
Using the following operators, write expressions and assign their values to variables:
g.==
h.<=
i.>=
j.!=
k.<
l.>
Now print each of the variables.
*/
package main

import "fmt"

func main() {
	a := (2 == 3)
	b := (6 <= 8)
	c := (8 >= 2)
	d := (6 != 3)
	e := (5 < 10)
	f := (67 > 10)

	fmt.Println(a, b, c, d, e, f)
}
