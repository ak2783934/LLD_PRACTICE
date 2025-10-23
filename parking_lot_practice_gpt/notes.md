1. I understood that we need to deisng a multi floor parking lot. 
the parking lot needs to have multiple type of vehicle allowed and only those vehicles will be allowed in such slots. 

- assuming that only one entry and exit is there. 
- in each floor, the slot with lower value is the nearest. like slot1, slot2, so slot1 is nearer than slot2. and will select this slot. 
- while entry of vehicle only, we create a ticket with start time. 
- assuming pricing will be implemented later. 



Class Design

// not actually need for this implementation. 
Vehicle{
    registration_number
    color 
}

type SlotType string
const (
    TwoWheeler SlotType = "TwoWheeler"
    FourWheeler SlotType = "FourWheeler"
    HeavyVehicle SlotType = "HeavyVehicle"
)

type Slot{
    floor string
    slotID string 
    slotType SlotType
    isOccupied bool 
}

ParkingLot{
    floors int
    slots map[int][]*Slot // sorted based on near -> floor -> []slots
    Tickets map[string]*Ticket
}

Ticket {
    ticketID string
    vehilceRegistrationNumber string 
    vehicleColor string 
    slotID string
    entryTime time.Time
    exitTime time.Time
}

CreateParkingLot(numFloors int, slotsPerFloor int) *ParkingLot{}
(p *ParkingLot)CreateTicket(regd string, vehicleColor string)*Ticket{}
(p *ParkingLot)ParkVehicle(vehicleType string, redgNum string, color string) (string,error)
(p *ParkingLot)UnParkVehicle(ticketID string) bool 
(p *ParkingLot)DisplayFreeSlot(vehicleType) []*Slots 
(p *ParkingLot)DisplayOccupiedSlots(vehicleType) []*Slot
