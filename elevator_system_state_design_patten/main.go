package main

func main() {
	elevator := &Elevator{floor: 0}
	elevator.setState(&IdleState{elevator})

	// Simulate Elevator Operations
	elevator.Request(5)
	elevator.Request(2)
	elevator.OpenDoor()
	elevator.CloseDoor()
	elevator.Request(0)
}
