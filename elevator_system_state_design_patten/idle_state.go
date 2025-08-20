package main

import "fmt"

type IdleState struct {
	ev *Elevator
}

func (I *IdleState) Request(floor int) {
	if floor > I.ev.floor {
		fmt.Printf("Elevator starting to move UP from floor %d to floor %d\n", I.ev.floor, floor)
		I.ev.setState(&MovingUpState{I.ev})
		I.ev.Request(floor)
	} else if floor < I.ev.floor {
		fmt.Printf("Elevator starting to move DOWN from floor %d to floor %d\n", I.ev.floor, floor)
		I.ev.setState(&MovingDownState{I.ev})
		I.ev.Request(floor)
	} else {
		fmt.Println("Elevator already at requested floor.")
		I.ev.setState(&DoorOpenState{I.ev})
		I.ev.OpenDoor()
	}
}

func (I *IdleState) OpenDoor() {
	fmt.Println("opening the door")
	I.ev.setState(&DoorOpenState{I.ev})
}

func (I *IdleState) CloseDoor() {
	fmt.Println("Door already closed in idle state.")
}
