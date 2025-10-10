package main

import (
	"fmt"
	"sort"
)

type Elevator struct {
	currentFloor  int
	maxFloor      int
	upQueue       []int
	downQueue     []int
	elevatorState ElevatorState
}

func NewElevator(maxFloor int) *Elevator {
	return &Elevator{
		currentFloor:  0,
		maxFloor:      maxFloor,
		upQueue:       make([]int, 0),
		downQueue:     make([]int, 0),
		elevatorState: &IdleState{},
	}
}

func (e *Elevator) EnqueFloor(floor int) {
	if floor < 0 || floor > e.maxFloor {
		fmt.Println("enter valid floor number")
		return
	}
	if e.currentFloor > floor {
		e.downQueue = append(e.downQueue, floor)
		sort.Slice(e.downQueue, func(a, b int) bool {
			return a > b
		})
	} else if e.currentFloor < floor {
		e.upQueue = append(e.upQueue, floor)
		sort.Slice(e.upQueue, func(a, b int) bool {
			return a < b
		})
	} else {
		fmt.Println("already in the same floor")
	}
}
func (e *Elevator) SetFloor(floor int) {
	e.currentFloor = floor
}
func (e *Elevator) GetFloor() int {
	return e.currentFloor
}
func (e *Elevator) SetElevatorState(elevatorState ElevatorState) {
	e.elevatorState = elevatorState
}
func (e *Elevator) GetElevatorState() ElevatorState {
	return e.elevatorState
}
func (e *Elevator) Request(floor int) {
	e.elevatorState.Request(e, floor)
}
func (e *Elevator) Move() {
	e.elevatorState.Move(e)
}
func (e *Elevator) OpenDoor() {
	e.elevatorState.OpenDoor(e)
}

type ElevatorState interface {
	Request(elevator *Elevator, floor int)
	OpenDoor(elevator *Elevator)
	Move(elevator *Elevator)
}

type IdleState struct{}

func (I *IdleState) Request(elevator *Elevator, floor int) {
	// here we need to start moving
	elevator.EnqueFloor(floor)
	if elevator.GetFloor() == floor {
		fmt.Println("Already in the same floor")
		elevator.SetElevatorState(&OpenDoorState{})
		elevator.OpenDoor()
		return
	}
	if elevator.GetFloor() > floor {
		elevator.SetElevatorState(&MovingDownState{})
	} else if elevator.GetFloor() < floor {
		elevator.SetElevatorState(&MovingUPState{})
	}
	elevator.Move()
}
func (I *IdleState) OpenDoor(elevator *Elevator) {
	fmt.Println("opening the doors")
	elevator.SetElevatorState(&OpenDoorState{})
}
func (I *IdleState) CloseDoor(elevator *Elevator) {
	fmt.Println("doors are closed")
}
func (I *IdleState) Move(elevator *Elevator) {
	fmt.Println("Elevator can't move, since its in idle state")
}

type MovingUPState struct{}

func (M *MovingUPState) Request(elevator *Elevator, floor int) {
	elevator.EnqueFloor(floor)
}
func (M *MovingUPState) OpenDoor(elevator *Elevator) {
	fmt.Println("can't open door, elevator moving ups")
}
func (M *MovingUPState) CloseDoor(elevator *Elevator) {
	fmt.Println("door already closed")
}
func (M *MovingUPState) Move(elevator *Elevator) {
	for {
		if len(elevator.upQueue) > 0 {
			target := elevator.upQueue[0]
			for elevator.GetFloor() < target {
				elevator.SetFloor(elevator.GetFloor() + 1)
				fmt.Println("moving up +1, current floor", elevator.GetFloor())
			}
			elevator.SetElevatorState(&OpenDoorState{})
			elevator.OpenDoor()
			elevator.upQueue = elevator.upQueue[1:]
			if len(elevator.upQueue) > 0 {
				fmt.Println("started moving up again")
				elevator.SetElevatorState(&MovingUPState{})
			} else {
				if len(elevator.downQueue) > 0 {
					fmt.Println("started moving down")
					elevator.SetElevatorState(&MovingDownState{})
					elevator.Move()
					return
				} else {
					fmt.Println("Set at idle state")
					elevator.SetElevatorState(&IdleState{})
					return
				}
			}
		}
	}
}

type MovingDownState struct{}

func (M *MovingDownState) Request(elevator *Elevator, floor int) {
	elevator.EnqueFloor(floor)
}
func (M *MovingDownState) OpenDoor(elevator *Elevator) {
	fmt.Println("can't open door, elevator moving")
}
func (M *MovingDownState) CloseDoor(elevator *Elevator) {
	fmt.Println("door already closed")
}
func (M *MovingDownState) Move(elevator *Elevator) {
	for {
		if len(elevator.downQueue) > 0 {
			target := elevator.downQueue[0]
			for elevator.GetFloor() > target {
				elevator.SetFloor(elevator.GetFloor() - 1)
				fmt.Println("moving down -1 current floor", elevator.GetFloor())
			}
			elevator.SetElevatorState(&OpenDoorState{})
			elevator.OpenDoor()
			elevator.downQueue = elevator.downQueue[1:]
			if len(elevator.downQueue) > 0 {
				fmt.Println("started moving down again")
				elevator.SetElevatorState(&MovingDownState{})
			} else {
				if len(elevator.upQueue) > 0 {
					fmt.Println("started moving up")
					elevator.SetElevatorState(&MovingUPState{})
					elevator.Move()
					return
				} else {
					fmt.Println("Set at idle state")
					elevator.SetElevatorState(&IdleState{})
					return
				}
			}
		}
	}
}

type OpenDoorState struct{}

func (O *OpenDoorState) Request(elevator *Elevator, floor int) {
	elevator.EnqueFloor(floor)
	if elevator.currentFloor > floor {
		elevator.SetElevatorState(&MovingDownState{})
		elevator.Move()
	} else if elevator.currentFloor < floor {
		elevator.SetElevatorState(&MovingDownState{})
		elevator.Move()
	} else {
		fmt.Println("aleady in same floor")
	}
}
func (O *OpenDoorState) OpenDoor(elevator *Elevator) {
	fmt.Println("Opening the door")
}
func (O *OpenDoorState) CloseDoor(elevator *Elevator) {
	fmt.Println("door already open")
}

func (O *OpenDoorState) Move(elevator *Elevator) {
	fmt.Println("can't move, doors open")
}

func main() {
	fmt.Println("starting elevator demo")
	newElevator := NewElevator(10)

	newElevator.Request(2)
	newElevator.Request(5)
	newElevator.Request(7)
	newElevator.Request(6)
	newElevator.Request(2)
	newElevator.Request(10)
}
