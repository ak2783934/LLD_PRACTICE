package main

import (
	"errors"
	"time"
)

type ParkingLot struct {
	ParkingSlots   []ParkingSlot
	Tickets        map[string]*Ticket
	ParkedVehicles map[string]*ParkingSlot
}

const (
	BIKE_RATE_PER_HOUR  = 10
	CAR_RATE_PER_HOUR   = 20
	TRUCK_RATE_PER_HOUR = 30
)

func (p *ParkingLot) FindSpot(vehicleType VehicleType) *ParkingSlot {
	for _, slot := range p.ParkingSlots {
		if slot.IsEmpty {
			switch slot.SlotType {
			case BIG:
				return &slot
			case MEDIUM:
				if vehicleType == CAR || vehicleType == BIKE {
					return &slot
				}
			case SMALL:
				if vehicleType == BIKE {
					return &slot
				}
			}
			return nil
		}
	}
	return nil
}

func (p *ParkingLot) ParkVehicle(vehicle *Vehicle) bool {
	parkingSpot := p.FindSpot(vehicle.VehicleType)

	ok := parkingSpot.Park(vehicle)
	if !ok {
		return false
	}

	// create a ticket
	p.Tickets[vehicle.RegNo] = &Ticket{
		StartTime: time.Now(),
		Vehicle:   *vehicle,
		Status:    UNPAID,
	}

	p.ParkedVehicles[vehicle.RegNo] = parkingSpot

	return true
}

func (p *ParkingLot) UnParkVehicle(vehicle Vehicle) (*Ticket, error) {
	parkingSpot := p.ParkedVehicles[vehicle.RegNo]

	if parkingSpot == nil {
		return nil, errors.New("vehicle not parked")
	}

	parkingSpot.UnPark()
	parkingTicket := p.Tickets[vehicle.RegNo]
	parkingTicket.EndTime = time.Now()

	parkingTicket.Price = caculatePrice(parkingTicket)
	parkingTicket.Status = PAID

	return parkingTicket, nil
}

func caculatePrice(ticket *Ticket) int {
	duration := ticket.EndTime.Sub(ticket.StartTime)

	switch ticket.Vehicle.VehicleType {
	case BIKE:
		return int(duration.Hours()) * BIKE_RATE_PER_HOUR
	case CAR:
		return int(duration.Hours()) * CAR_RATE_PER_HOUR
	case TRUCK:
		return int(duration.Hours()) * TRUCK_RATE_PER_HOUR
	}
	return 0
}
