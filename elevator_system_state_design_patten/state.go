package main

type ElevatorState interface {
	Request(floor int)
	OpenDoor()
	CloseDoor()
}

/*
Different states

IdleState -
MovingUpState -
MovingDownState -
DoorOpenState - 
MaintenanceState


*/
