package main

import (
	"errors"
	"time"
)

type VehicleType string

const (
	Bike  VehicleType = "BIKE"
	Car   VehicleType = "CAR"
	Truck VehicleType = "TRUCK"
)

type Vehichle struct {
	RedgNo string
	Type   VehicleType
}

type ParkingSpotType string

const (
	TwoWheeler  ParkingSpotType = "TWO_WHEELAR"
	FourWheeler ParkingSpotType = "FOUR_WHEELAR"
	HEAVY       ParkingSpotType = "HEAVY"
)

type ParkingSpot struct {
	Id                  int // incremental based on distance
	floor               int
	IsVacant            bool
	Type                ParkingSpotType
	ParkedVehicleRegdNo string
}

type TicketStatus string

const (
	Unpaid TicketStatus = "UNPAID"
	Paid   TicketStatus = "PAID"
)

type Ticket struct {
	Id            int
	VehicleRegdNo string
	ParkingSpotID int
	EntryTime     time.Time
	Status        TicketStatus
}

type ParkingLot struct {
	Floors               int
	ParkingSpotsPerFloor []map[string]*ParkingSpot
	ParkingTickets       []*Ticket
	pricingStrategy      PricingStrategy
	parkingStrategy      ParkingStrategy
}

func (p *ParkingLot) ParkVehicle(vehicle Vehichle) (int, error) {
	// find the spot using the strategy
	spot, err := p.parkingStrategy.findParkingSpot(vehicle)
	if err != nil {
		return 0, errors.New("parking spot not found")
	}

	spot.IsVacant = false
	spot.ParkedVehicleRegdNo = vehicle.RedgNo
	// create ticket and return ticket id.
	ticket := Ticket{
		Id:            1213,
		VehicleRegdNo: vehicle.RedgNo,
		ParkingSpotID: spot.Id,
		EntryTime:     time.Now(),
		Status:        Unpaid,
	}
	p.ParkingTickets = append(p.ParkingTickets, &ticket)

	return ticket.Id, nil
}

func (p *ParkingLot) UnParkVehicle(vehicle Vehichle) (int, error) {
	// find the spot for this vehicle?
	for _, floorSpots := range p.ParkingSpotsPerFloor {
		for _, spot := range floorSpots {
			if spot.ParkedVehicleRegdNo == vehicle.RedgNo {
				// make the spot vacent.
				spot.IsVacant = true
				spot.ParkedVehicleRegdNo = ""

				// find the ticket for this vehicle
				var parkingTicket *Ticket
				for _, ticket := range p.ParkingTickets {
					if ticket.Status == Unpaid && ticket.VehicleRegdNo == vehicle.RedgNo {
						parkingTicket = ticket
					}
				}
				if parkingTicket == nil {
					return 0, errors.New("ticket not foudn for given vehicle")
				}

				// calculate the price using pricing strategy
				cost := p.pricingStrategy.Calculate(parkingTicket)
				return cost, nil
			}
		}
	}

	return 0, errors.New("Given vehicle is not found in parking spot")
}
