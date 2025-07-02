package main

type VehicleType int

const (
	BIKE VehicleType = iota
	CAR
	TRUCK
)

type Vehicle struct {
	RegNo       string
	VehicleType VehicleType
}
