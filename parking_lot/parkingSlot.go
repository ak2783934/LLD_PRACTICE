package main

type ParkigSlotType int

const (
	SMALL ParkigSlotType = iota
	MEDIUM
	BIG
)

type ParkingLocation struct {
	FloorNumber int
	SpotNumber  string
}

type ParkingSlot struct {
	SlotType        ParkigSlotType
	ParkingLocation ParkingLocation
	IsEmpty         bool
	ParkedVehicle   *Vehicle
}

func CreateParkingSlot(slotType ParkigSlotType, location ParkingLocation) *ParkingSlot {
	return &ParkingSlot{
		SlotType:        slotType,
		ParkingLocation: location,
	}
}

func (p *ParkingSlot) Park(vehicle *Vehicle) bool {
	if p.IsEmpty {
		if p.SlotType == SMALL && vehicle.VehicleType == BIKE {
			p.ParkedVehicle = vehicle
			p.IsEmpty = false
			return true
		}

		if p.SlotType == MEDIUM && (vehicle.VehicleType == CAR || vehicle.VehicleType == BIKE) {
			p.ParkedVehicle = vehicle
			p.IsEmpty = false
			return true
		}

		if p.SlotType == BIG {
			p.ParkedVehicle = vehicle
			p.IsEmpty = false
			return true
		}

		return false
	} else {
		return false
	}
}

func (p *ParkingSlot) UnPark() bool {
	if !p.IsEmpty {
		p.IsEmpty = true
		p.ParkedVehicle = nil
		return true
	} else {
		return false
	}
}
