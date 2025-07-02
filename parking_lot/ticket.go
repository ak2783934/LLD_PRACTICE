package main

import "time"

type TicketStatus int

const (
	UNPAID TicketStatus = iota
	PAID
)

type Ticket struct {
	StartTime time.Time
	EndTime   time.Time
	Vehicle   Vehicle
	Price     int
	Status    TicketStatus
}
