/*Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
“defer” keyword used to delay the execution of a function or a statement until the nearby function returns
*/
package main

import "fmt"

func main() {
	defer fmt.Println("world1")
	defer fmt.Println("world2")
	defer fmt.Println("world3")
	fmt.Println("hello")
}
