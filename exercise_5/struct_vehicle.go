/*
●Create a new type: vehicle.
    ○The underlying type is a struct.
    ○The fields:
        ■doors
        ■color
●Create two new types: truck & sedan.
    ○The underlying type of each of these new types is a struct.
    ○Embed the “vehicle” type in both truck & sedan.
    ○Give truck the field “fourWheel” which will be set to bool.
    ○Give sedan the field “luxury” which will be set to bool. solution
●Using the vehicle, truck, and sedan structs:
    ○using a composite literal, create a value of type truck and assign values to thefields;
    ○using a composite literal, create a value of type sedan and assign values to thefields.
●Print out each of these values.●Print out a single field from each of these values.
*/

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type fourwheel struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	echoSport := fourwheel{
		vehicle: vehicle{
			doors: 2,
			color: "Green",
		},
		fourWheel: true,
	}

	sed := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(echoSport.color)
	fmt.Println(sed)
}
