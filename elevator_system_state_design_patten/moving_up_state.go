package main

import "fmt"

type MovingUpState struct {
	ev *Elevator
}

func (M *MovingUpState) Request(floor int) {
	fmt.Printf("Elevator moving UP... reached floor %d\n", floor)
	M.ev.floor = floor
	M.ev.setState(&DoorOpenState{M.ev})
	M.ev.OpenDoor()
}

func (M *MovingUpState) OpenDoor() {
	fmt.Println("Cannot open doors while moving UP.")
}

func (M *MovingUpState) CloseDoor() {
	fmt.Println("Door already closed in moving up state.")
}
