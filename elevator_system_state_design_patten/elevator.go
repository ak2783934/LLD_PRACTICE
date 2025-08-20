package main

type Elevator struct {
	floor int
	state ElevatorState
}

func (e *Elevator) setState(state ElevatorState) {
e.state = state
}

func (e *Elevator) Request(floor int) {
	e.state.Request(floor)
}

func (e *Elevator) OpenDoor() {
	e.state.OpenDoor()
}

func (e *Elevator) CloseDoor() {
	e.state.CloseDoor()
}
