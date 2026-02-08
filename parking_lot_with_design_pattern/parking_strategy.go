package main

type ParkingStrategy interface {
	findParkingSpot(Vehichle) (*ParkingSpot, error)
}

type NearestParkingStrategy struct {
}

func (n *NearestParkingStrategy) findParkingSpot(v Vehichle) (*ParkingSpot, error) {
	return nil, nil
}
