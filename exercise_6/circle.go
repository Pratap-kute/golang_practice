/*
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
	○ square area = L * W
●create a type SHAPE that defines an interface as anything that has the AREA method
●create a func INFO which takes type shape and then prints the area

●create a value of type square
●create a value of type circle
●use func info to print the area of square
●use func info to print the area of circle

*/

/*
two struct types, square and circle, each with their respective properties.
It also defines an interface type shape, which has a single method area() that returns a float64 value.
The square and circle types both implement the shape interface by defining their own area() 
methods that calculate the area of the respective shape.

The main() function creates an instance of a square and a circle struct, and passes them to the info() function.
The info() function takes a shape interface as an argument, and calls the area() 
method on the argument, which calculates the area of the shape and prints it to the console using the fmt package.

The output of this program will be the area of the square and the circle, respectively,
which is calculated using the area() method of each shape. 
The area of the square is calculated as the product of the length and width properties, 
while the area of the circle is calculated as π times the square of the radius.

*/

package main

import "fmt"

func main() {
	sq := square{
		length: 10,
		width:  10,
	}
	ci := circle{
		radius: 10,
	}
	info(sq)
	info(ci)
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}
