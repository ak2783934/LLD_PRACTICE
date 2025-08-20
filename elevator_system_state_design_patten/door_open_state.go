package main

import "fmt"

type DoorOpenState struct {
	ev *Elevator
}

func (D *DoorOpenState) Request(floor int) {
	fmt.Println("can not move lift as door is open, Closing doors first")
	D.ev.CloseDoor()
	D.ev.Request(floor)
}

func (D *DoorOpenState) OpenDoor() {
	fmt.Println("Doors are already open.")
}

func (D *DoorOpenState) CloseDoor() {
	fmt.Println("Closing the doors")
	D.ev.setState(&IdleState{D.ev})
}
