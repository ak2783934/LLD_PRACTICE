package main

import (
	"errors"
	"time"
)

// Entities

type BusType string

const (
	Sleeper     BusType = "SLEEPER"
	Seater      BusType = "SEATER"
	SemiSleeper BusType = "SEMI_SLEEPER"
)

type SeatType string

const (
	SleeperSeatType     SeatType = "SLEEPER"
	SeaterSeatType      SeatType = "SEATER"
	SemiSleeperSeatType SeatType = "SEMI_SLEEPER"
)

type BusSeatLayout struct {
	Row      int
	Cols     int
	SeatType SeatType
}

type Bus struct {
	BusID         string
	RegdNo        string
	OperatorName  string
	BusType       BusType
	TotalSeat     int
	BusSeatLayout []BusSeatLayout
}

type SeatStatus string

const (
	Booked    SeatStatus = "BOOKED"
	Held      SeatStatus = "HELD"
	Available SeatStatus = "AVAILABLE"
)

type BusSeat struct {
	SeatID        int
	SeatStatus    SeatStatus
	heldTimeStamp time.Time
	SeatType      SeatType
	Cost          int
}

func (S *BusSeat) isSeatAvailable() bool {
	switch S.SeatStatus {
	case Held:
		elapsedTime := time.Since(S.heldTimeStamp)
		if elapsedTime > time.Minute*5 {
			// relase the lock
			return true
		} else {
			return false
		}
	case Available:
		return true
	case Booked:
		return false
	default:
		return false
	}
}

func (seat *BusSeat) BookSeat() {
	seat.SeatStatus = Held
	seat.heldTimeStamp = time.Now()
}
func (seat *BusSeat) ConfirmSeat() {
	seat.SeatStatus = Booked
	seat.heldTimeStamp = time.Time{}
}
func (seat *BusSeat) CancelSeat() {
	seat.SeatStatus = Available
	seat.heldTimeStamp = time.Time{}
}

type BusTrip struct {
	BusTripID         string
	StartPoint        string
	EndPoint          string
	IntermediateStops []string // ordered in the order of arrival
	StartDateTime     time.Time
	EndDateTime       time.Time
	Seats             []*BusSeat
}

func (B *BusTrip) GetAvailableSeats() []*BusSeat {
	availableSeats := []*BusSeat{}
	for _, seat := range B.Seats {
		if seat.isSeatAvailable() == true {
			availableSeats = append(availableSeats, seat)
		}
	}
	return availableSeats
}

type TicketStatus string

const (
	PendingPayment   TicketStatus = "PENDING_PAYMENT"
	ConfirmPayment   TicketStatus = "CONFIRM_PAYMENT"
	CancelledPayment TicketStatus = "CANCELLED_PAYMENT"
)

type BusTicket struct {
	TicketID     string
	TripID       string
	UserID       string
	Seats        []*BusSeat
	Amount       int
	TicketStatus TicketStatus
}

type BusAggregator struct {
	Buses map[string]*Bus
	Trips   []*BusTrip
	Tickets map[string]*BusTicket
}

func NewBusAggregator() *BusAggregator {
	// singleton can be use here. 
	return &BusAggregator{
		Trips:   make([]*BusTrip, 0),
		Tickets: make(map[string]*BusTicket),
	}
}

func (B *BusAggregator) SearchBus() []*BusTrip {
	// extensive logic to match based on filter. will implement it later.

	// it will be based on matching the destination, source and intermediate details. 
	// also based on filter bus type, timeing and other details. 
	return nil
}

func (B *BusAggregator) SelectSeats(userID string, trip *BusTrip, seats []*BusSeat) (string, error) {
	// validate if the given seats exist?
	availableSeats := trip.GetAvailableSeats()

	for _, seat := range seats {
		if seat.isSeatAvailable() == false {
			return "", errors.New("seat already booked")
		}
		isFound := false
		for _, availableSeat := range availableSeats {
			if seat.SeatID == availableSeat.SeatID {
				isFound = true
			}
		}
		if isFound == false {
			return "", errors.New("seat doesn't belong to bus")
		}
	}

	// mark them blocked.
	cost := 0
	for _, seat := range seats {
		seat.BookSeat()
		cost += seat.Cost
	}

	// create ticket with price.
	ticket := &BusTicket{
		TicketID:     GenerateEightCharID(),
		TripID:       trip.BusTripID,
		UserID:       userID,
		TicketStatus: PendingPayment,
		Amount:       cost,
		Seats:        seats,
	}
	B.Tickets[ticket.TicketID] = ticket
	return ticket.TicketID, nil
}

func (B *BusAggregator) ConfirmTicket(ticketID string) error {
	ticket, ok := B.Tickets[ticketID]
	if !ok {
		return errors.New("invalid ticket id")
	}
	seats := ticket.Seats
	for _, seat := range seats {
		seat.ConfirmSeat()
	}
	ticket.TicketStatus = ConfirmPayment
	return nil
}

func (B *BusAggregator) CancelTicket(ticketID string) error {
	ticket, ok := B.Tickets[ticketID]
	if !ok {
		return errors.New("invalid ticket id")
	}
	seats := ticket.Seats
	for _, seat := range seats {
		seat.CancelSeat()
	}
	ticket.TicketStatus = CancelledPayment
	return nil
}

// interfaces

DynamicPricingStrategy



