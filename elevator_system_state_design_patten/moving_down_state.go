package main

import "fmt"

type MovingDownState struct {
	ev *Elevator
}

func (M *MovingDownState) Request(floor int) {
	fmt.Printf("Elevator moving Down... reached floor %d\n", floor)
	M.ev.floor = floor
	M.ev.setState(&DoorOpenState{M.ev})
	M.ev.OpenDoor()
}

func (M *MovingDownState) OpenDoor() {
	fmt.Println("Cannot open doors while moving Down.")
}

func (M *MovingDownState) CloseDoor() {
	fmt.Println("Door already closed in moving down state.")
}
