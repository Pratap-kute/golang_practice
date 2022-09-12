package main
import "fmt"
// iota increment the next veriable by one 
const (
	a=iota
	b
	c
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
}