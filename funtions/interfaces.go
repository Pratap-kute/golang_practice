// Description: Interfaces
/*
This code is a Go program that defines two structs - person and secretAgent. 
The secretAgent struct embeds the person struct, meaning it inherits all of its fields and methods.

The program then defines two methods - speak() for person and secretAgent. 
The speak() method for person simply prints out the person's name with a message indicating that the person is speaking. 
The speak() method for secretAgent does the same thing, but with a message indicating that the secret agent is speaking.

Next, the program defines an interface human that specifies a single method - speak(). 
This interface is implemented by both the person and secretAgent structs.

Finally, in the main() function, the program creates two instances of secretAgent and one instance of person. 
It then prints out the sa1 instance, calls the speak() method on both sa1 and sa2, prints out the p1 instance, 
and calls the bar() function with all three instances as arguments.

The bar() function takes a parameter of type human, which means it can accept any object that implements the speak() method. 
This function simply prints out a message indicating that the object was passed into bar().
*/

package main

import "fmt"

type person struct {
    first string
    last  string
}

type secretAgent struct {
    person
    ltk bool
}

func (p person) speak() {
    fmt.Println("I am", p.first, p.last,"- the person speaks")
}

func (sa secretAgent) speak() {
    fmt.Println("I am", sa.first, sa.last, "- the secret agent speaks")
}

func bar(h human) {
    fmt.Println("I was passed into bar", h)
}

type human interface {
    speak()
}

func main() {
    sa1 := secretAgent{
        person: person{
            first: "James",
            last:  "Bond",
        },
        ltk: true,
    }

    sa2 := secretAgent{
        person: person{
            first: "Miss",
            last:  "Moneypenny",
        },
        ltk: true,
    }

    p1 := person{
        first: "Miss",
        last:  "Moneypenny",
    }

    fmt.Println(sa1)
    sa1.speak()
    sa2.speak()

    fmt.Println(p1)
    bar(sa1)
    bar(sa2)
    bar(p1)
}
