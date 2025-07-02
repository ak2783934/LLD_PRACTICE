package main

import "fmt"

type Direction int

const (
	IDLE Direction = iota
	UP
	DOWN
)

type DoorStatus int

const (
	OPEN DoorStatus = iota
	CLOSED
)

type Request struct {
	IsExternal bool
	Direction  Direction
	Floor      int
}

type Elevator struct {
	ID           int
	CurrentFloor int
	TotalFloor   int
	Requests     []Request
	DoorStatus   DoorStatus
	Direction    Direction
}

func (e *Elevator) AddRequest(request Request) {
	e.Requests = append(e.Requests, request)
	fmt.Printf("Elevator %d added request for floor %d\n", e.ID, request.Floor)
}

func (e *Elevator) Step() {
	if len(e.Requests) == 0 {
		e.Direction = IDLE
		return
	}

	requestedFloor := e.Requests[0].Floor
	if requestedFloor > e.CurrentFloor {
		e.CurrentFloor++
		e.Direction = UP
	} else if requestedFloor < e.CurrentFloor {
		e.CurrentFloor--
		e.Direction = DOWN
	} else {
		e.DoorStatus = OPEN

		fmt.Print("reached the floor")
		e.DoorStatus = CLOSED
		e.Requests = e.Requests[1:]
	}
}

type ElevatorController struct {
	Elevators []*Elevator
}

func (c *ElevatorController) HandleExternalRequest(req Request) {
	var chosen *Elevator
	for _, e := range c.Elevators {
		if e.Direction == IDLE {
			chosen = e
			break
		}
	}
	if chosen == nil {
		chosen = c.Elevators[0]
	}
	chosen.AddRequest(req)
}

func (c *ElevatorController) StepAll() {
	for _, e := range c.Elevators {
		e.Step()
	}
}
